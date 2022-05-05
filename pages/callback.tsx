import axios from 'axios'
import { NOTIMP } from 'dns'
import type { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { LineLoginUrl, LINE_CLIENT_ID, LINE_CLIENT_SECRET, LINE_TOKEN_URL } from '../lib/param'
import styles from '../styles/Home.module.css'


const CallbackPage: NextPage = () => {
  const router = useRouter();
  const code = router.query.code;

  if (typeof(code) !== 'string') {
    return <div>Loading...</div>
  }

  (async  () => {
    const resp = await axios.post('/api/check_auth', {
      code: code
    });

    console.log(resp)
  })()


  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>{code}</p>
      </main>
    </div>
  )
}

export default CallbackPage
