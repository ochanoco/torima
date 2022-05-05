const LINE_REDIRECT_URI = process.env.LINE_REDIRECT_URI as string;
const LINE_CLIENT_ID = process.env.LINE_CLIENT_ID as string;
const LINE_CLIENT_SECRET = process.env.LINE_CLIENT_SECRET as string;
const NEXT_PUBLIC_LINE_LOGIN_URL = process.env.NEXT_PUBLIC_LINE_LOGIN_URL as string;

const LineLoginUrl = (state: string, nonce: string) => {
    let url = NEXT_PUBLIC_LINE_LOGIN_URL;
    url = url.replace("{{LINE_REDIRECT_URI}}", LINE_REDIRECT_URI)
    url = url.replace("{{LINE_CLIENT_ID}}", LINE_CLIENT_ID)
    
    url = url.replace("{{state}}", state)
    url = url.replace("{{nonce}}", nonce)

    return url
}

const LINE_TOKEN_URL = 'https://api.line.me/oauth2/v2.1/token'


export {
    LINE_CLIENT_ID,
    LINE_CLIENT_SECRET,
    LINE_TOKEN_URL,
    LINE_REDIRECT_URI,
    LineLoginUrl,
}