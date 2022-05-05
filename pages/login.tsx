import type { NextPage } from 'next'
import Head from 'next/head'
import Link from 'next/link'
import React from 'react'
import styles from '../styles/Home.module.css'

const LoginPage: NextPage = () => {
  const LINE_LOGIN_URL = process.env.NEXT_PUBLIC_LINE_LOGIN_URL as string

  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Link href={LINE_LOGIN_URL}>
          <a> Login </a>
        </Link>
      </main>
    </div>
  )
}

export default LoginPage
