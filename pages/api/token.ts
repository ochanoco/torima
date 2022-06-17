import { getSession } from "next-auth/react";
import { NextApiRequest, NextApiResponse } from "next/types";

import { generateToken } from "../../lib/jwt";

export default async function handler(req: NextApiRequest, res: NextApiResponse): Promise<void> {
    const session = await getSession()

    console.log("session: ", session)

    if (session?.token) {
        const token = generateToken(session?.token.sub)

        res.status(200).json({
            token
        })

    }
    else
        res.status(400).json({})
}