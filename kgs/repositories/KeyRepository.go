package repositories
import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
	"fmt"
	"log"
)

type Key struct {
	Value string 
	IsUsed bool 
}
type KeyRepository interface {
	GetUnusedKey() (string, error)
	InsertKey(key string) (string, error)
	Exists(key string) (bool, error)
}

type MongoKeyRepository struct {
	// keys map[string]bool
	// unusedKeys []string
	*mongo.Collection
}

func NewMongoKeyRepository() *MongoKeyRepository {
	databaseName := utils.Getenv("DB_NAME")
	databaseHost := utils.Getenv("DB_HOST")
	databasePort := utils.Getenv("DB_PORT")

	uri := "mongodb://"+databaseHost + ":" + databasePort
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Mongo connection is not successful")
		panic(err)
	}
	db := client.Database(databaseName)

	fmt.Println("Mongo connection is successful")

	repo := &MongoKeyRepository{db.Collection("key")}

	return repo
}

func (repo *MongoKeyRepository) GetUnusedKey() (string, error) {
	var key Key
	err := repo.Collection.FindOne(context.Background(), bson.M{"isUsed": false}).Decode(&key)
	if err != nil {
		log.Fatal(err)
	} else {
		if _, err = repo.Collection.UpdateOne(context.Background(), bson.M{"value": key.Value}, bson.M{"$set": bson.M{"isUsed": true}}); err != nil {
			log.Fatal(err)
		}
	}
	return key.Value, nil
}

func (repo *MongoKeyRepository) InsertKey(key string) (string, error) {
	exists, _ := repo.Exists(key)
	if (!exists) {
		_, err := repo.Collection.InsertOne(context.Background(), bson.M{"value": key, "isUsed": false})
		if err != nil {
			return "", err
		}

		return "success", nil
	}
	return "", nil
}

func (repo *MongoKeyRepository) Exists(key string) (bool, error) {
	result := repo.Collection.FindOne(context.Background(), bson.M{"value": key})
	var k Key
	err := result.Decode(&k)
	if err != nil || k.IsUsed {
		// NOTE: key not found
		return false, err
	}
	return true, err
}

func (repo *MongoKeyRepository) FindAll() (error) {
	cur, err := repo.Collection.Find(context.Background(), bson.D{})
	// var k Key
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var k Key
		err := cur.Decode(&k)
		if err != nil { 
			log.Fatal(err) 
		}
		// do something with result...
		fmt.Println(cur)
		// To get the raw bson bytes use cursor.Current
		// raw := cur.Current
		// do something with raw...
	}
	if err = cur.Err(); err != nil {
		return err
	}
	return err
}