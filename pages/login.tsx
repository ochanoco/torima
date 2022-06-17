import type { NextPage } from 'next'
import Head from 'next/head'
import React from 'react'
import styles from '../styles/Home.module.css'
import { useSession, signIn, signOut } from "next-auth/react"
import { useRouter } from 'next/router'

import { generateToken } from '../lib/jwt'


const LoginPage: NextPage = () => {
  const { data: session } = useSession()
  const router = useRouter()

  if (session) {
    console.log("session", session)
    const token = generateToken('hello')

    const refererUrl = localStorage.getItem('referer')
    const redirectUrl = `${refererUrl}?token=${token}`

    localStorage.setItem('referer', redirectUrl as string)

    if (!refererUrl) return (
      <div>
        <div>Logged in</div>
        <br />
        <button onClick={() => signOut()}>Log out</button>
      </div>
    )

    router.push(refererUrl as string)
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Login Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <button onClick={() => signIn()}>Log in</button>
      </main>
    </div>
  )
}

export default LoginPage
