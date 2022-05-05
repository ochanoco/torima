// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import axios, { AxiosError } from 'axios';
import type { NextApiRequest, NextApiResponse } from 'next'
import { LINE_CLIENT_ID, LINE_CLIENT_SECRET, LINE_REDIRECT_URI, LINE_TOKEN_URL } from '../../lib/param';


interface NextCheckAuthReq extends NextApiRequest {
  body: {
    code: number;
  };
}


export default async function handler(
  req: NextCheckAuthReq,
  res: NextApiResponse<string>
) {
  const param = {
    code: req.body.code,
    grant_type: 'authorization_code',
    client_id: LINE_CLIENT_ID,
    client_secret: LINE_CLIENT_SECRET,
    redirect_uri: LINE_REDIRECT_URI
  };

  console.log(param)

  try {
    const resp = await axios.post(LINE_TOKEN_URL, param);
    res.status(200)
    res.end("ok")
  } catch (e : AxiosError) {
    console.log('---- Error ----')
    console.log(e.response.data)
    console.log(e.request)  
    // console.log(e.response)
    res.status(400)
    res.end(e.response.statusText)
  }
}
