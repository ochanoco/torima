import { TEXT_ON_LOGIN_PAGE } from '../languages/ja_jp'
import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'


const LoginPage: NextPage = ({ }) => {
  return (
    <div className={styles.container}>
      <Head>
        <title>{TEXT_ON_LOGIN_PAGE.TITLE}</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <h1>{TEXT_ON_LOGIN_PAGE.HEADER}</h1>
        <p>{TEXT_ON_LOGIN_PAGE.MESSAGE}</p >
        <button onClick={() => {
          location.href = '/ochanoco/redirect'
        }}>{TEXT_ON_LOGIN_PAGE.BUTTON}</button>
      </main>
    </div>
  )
}


export default LoginPage
