import * as jwt from 'jsonwebtoken';
import { Secret, SignOptions } from 'jsonwebtoken';


const generateToken = (userId: string): string => {
    const jwtPayload = {
        userId,
    }

    const jwtSecret: Secret = 'process.env.JWT_SECRET_KEY' as string
    const jwtOptions: SignOptions = {
        algorithm: 'HS256',
        expiresIn: '10s',
        issuer: 'process.env.JWT_ISS' as string
    }

    const token = jwt.sign(jwtPayload, jwtSecret, jwtOptions)
    return token
}

export { generateToken }