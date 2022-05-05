import type { NextPage } from 'next'
import Head from 'next/head'
import React from 'react'
import styles from '../styles/Home.module.css'
import { useSession, signIn, signOut } from "next-auth/react"


const LoginPage: NextPage = () => {
  const { data: session } = useSession()

  if (session) {
    console.log("session", session)

    return <div>
        <div>Logged in</div>
        <br />
        <button onClick={() => signOut()}>Log out</button>
      </div>
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
