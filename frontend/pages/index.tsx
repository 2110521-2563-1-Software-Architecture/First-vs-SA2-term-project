import ShortenerPage from 'features/shortener/ShortenerPage'
import { Navbar } from 'components/Navbar'
import { ShortenHistory } from 'components/ShortenHistory'
import Head from 'next/head'

export default function Home() {
  return (
    <>
      <Navbar />
      <ShortenerPage />
      <ShortenHistory />
    </>
  )
}
