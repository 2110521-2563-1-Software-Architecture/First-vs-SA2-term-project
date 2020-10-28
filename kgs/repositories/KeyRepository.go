package repositories

type KeyRepository interface {
	GetUnusedKey() (string, error)
	InsertKey(key string) (string, error)
}

type MemoryKeyRepository struct {
	keys map[string]bool
	unusedKeys []string
}

func NewMemoryKeyRepository() *MemoryKeyRepository {
	repo := MemoryKeyRepository{}
	repo.keys = make(map[string]bool)
	repo.unusedKeys = make([]string, 0)
	return &repo
}

func (repo *MemoryKeyRepository) GetUnusedKey() (string, error) {
	var key = repo.unusedKeys[0]
	repo.keys[key] = true
	repo.unusedKeys = repo.unusedKeys[1:]
	return key, nil
}

func (repo *MemoryKeyRepository) InsertKey(key string) (string, error) {
	repo.unusedKeys = append(repo.unusedKeys, key)
	repo.keys[key] = false
	return key, nil
}