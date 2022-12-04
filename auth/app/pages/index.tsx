import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'

const MainPage: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>hello</p>
      </main>
    </div>
  )
}

export default MainPage
