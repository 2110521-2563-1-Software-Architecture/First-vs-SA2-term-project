import axios from 'axios'
import { API_ENDPOINT } from './config'

export const getShortenURL = async(originalURL: string) => {
    const resp = await (await axios.post(`${API_ENDPOINT}/shorten`, {url: originalURL})).data
    console.log(resp)
    return resp
}

export const redirectTo = async(url: string) => {
    const resp = await (await axios.get(`${API_ENDPOINT}/url`)).data
    console.log(resp)
}