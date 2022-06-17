import type { NextPage } from 'next'
import { useEffect } from 'react'
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import router, { useRouter } from 'next/router'

const REDIRECT_URL = "http://localhost:3000/login"

const RedirectPage: NextPage = () => {
  const router = useRouter();

  const { referer } = router.query;
  if (referer)
    localStorage.setItem('referer', referer as string)

  useEffect(() => {
    setTimeout(() => {
      location.href = REDIRECT_URL
    }, 3000)
  })
  return (
    <div className={styles.container}>
      <Head>
        <title>Redirect Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>We redirect to {REDIRECT_URL} after 3 sec.</p>
      </main>
    </div>
  )
}

export default RedirectPage
