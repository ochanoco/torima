import { NOTIMP } from 'dns'
import type { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import styles from '../styles/Home.module.css'

const CallbackPage: NextPage = () => {
  const router = useRouter();


  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>{router.query.code}</p>
      </main>
    </div>
  )
}

export default CallbackPage
