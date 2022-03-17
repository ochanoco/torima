import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'

const LoginPage: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>login</p>
      </main>
    </div>
  )
}

export default LoginPage
