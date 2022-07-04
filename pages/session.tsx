import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import nookies from 'nookies'

const SessionPage: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Redirect Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>Hello</p>
      </main>
    </div>
  )
}

export async function getServerSideProps(ctx: any) {
  const token = ctx.query.token;

  const cookies = nookies.get(ctx)

  nookies.set(ctx, 'ochanoco-token', token as string, {
    maxAge: 30 * 24 * 60 * 60,
    path: '/',
  })

  return { props: cookies }
}


export default SessionPage
