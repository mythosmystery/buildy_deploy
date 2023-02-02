import express, { Request, Response } from 'express'

import bodyParser from 'body-parser'

const app = express()

app.use(bodyParser.json())

app.get('/', (_req: Request, res: Response) => {
    res.send('Hello World!')
})

app.post('/webhooks', (req: Request, res: Response) => {
    console.log(req.body)
    res.send('OK')
})

app.listen(8001, () => {
    console.log('Server started on port 8001')
})