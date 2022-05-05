import type { NextPage } from 'next'
import Head from 'next/head'
import Link from 'next/link'
import React from 'react'
import { LineLoginUrl } from '../lib/param'
import styles from '../styles/Home.module.css'

const LoginPage: NextPage = () => {
  const lineLoginUrl = LineLoginUrl('10000', '10000')
  console.log(lineLoginUrl)

  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Link href={lineLoginUrl}>
          <a> Login </a>
        </Link>
      </main>
    </div>
  )
}

export default LoginPage
