import NextAuth from 'next-auth'
import LineProvider from "next-auth/providers/line";

export default NextAuth({
    providers: [
        LineProvider({
            clientId: process.env.LINE_CLIENT_ID,
            clientSecret: process.env.LINE_CLIENT_SECRET,
        })
    ],
    callbacks: {
        async session({ session, token, user }) {
            session.token = token

            // console.log({ session, token, user })
            return session
        }
    }
})