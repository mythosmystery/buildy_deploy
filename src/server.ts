import express, { Request, Response } from 'express'
import { PushEvent } from '@octokit/webhooks-types'

import bodyParser from 'body-parser'

const app = express()

app.use(bodyParser.json())

app.get('/', (_req: Request, res: Response) => {
    res.send('Hello World!')
})

app.post('/webhooks', (req, res: Response) => {
    const body: PushEvent = req.body
     if (body.ref === `refs/heads/${process.env.BRANCH || 'main'}`) {
        console.log('Main branch updated, redeploying...')
     }
    res.send('OK')
})

app.listen(8001, () => {
    console.log('Server started on port 8001')
})