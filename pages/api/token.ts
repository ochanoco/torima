import { getSession } from "next-auth/react";
import { NextApiRequest, NextApiResponse } from "next/types";

import { generateToken } from "../../lib/jwt";

export default async function handler(req: NextApiRequest, res: NextApiResponse): Promise<void> {
    const session = await getSession()
    const token = generateToken(session?.token.sub)

    res.status(200).json({
        token
    })
}