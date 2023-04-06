import { TEXT_ON_REDIRET_PAGE } from '../languages/ja_jp'
import { OCHANOCO_CONFIG } from '../config'
import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import { useEffect, useState } from 'react'


const RedirectPage: NextPage = ({ }) => {
    useEffect(() => {
        setTimeout(() => {
        }, 6000)
        location.href = `http://127.0.0.1:8081/login?client_id=${document.location.host}`
    }, [])

    return (
        <div className={styles.container}>
            <Head>
                <title>{TEXT_ON_REDIRET_PAGE.TITLE}</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <main>
                <h1>{TEXT_ON_REDIRET_PAGE.HEADER}</h1>
                <p>{TEXT_ON_REDIRET_PAGE.MESSAGE}</p>
            </main>
        </div>
    )
}

export default RedirectPage
