import type { NextPage } from 'next'
import Head from 'next/head'
import React, { useState } from 'react'
import styles from '../styles/Home.module.css'
import { useSession, signIn, signOut, getSession } from "next-auth/react"
import { generateToken } from '../lib/jwt'


const LoginPage: NextPage = ({ token }) => {
  const { data: session } = useSession()

  if (session) {
    console.log("data: ", token)

    const refererUrl = localStorage.getItem('referer')
    const redirectUrl = `${refererUrl}?token=${token}`


    if (!refererUrl) return (
      <div>
        <div>Logged in</div>
        <br />
        <button onClick={() => signOut()}>Log out</button>
      </div>
    )

    console.log("token: ", redirectUrl)
    location.href = redirectUrl
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


export async function getServerSideProps(context) {
  let token = ''
  const session = await getSession(context)

  if (session?.token) {
    token = generateToken(session?.token.sub)
  }

  console.log(token)

  return {
    props: {
      token
    },
  }
}


export default LoginPage
