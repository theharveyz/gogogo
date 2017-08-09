import path from 'path'
import express from 'express'
import serve from 'serve-static' // return async
import opn from 'opn'

const app = express()
const port = process.env.PORT || 1900

app.use(serve(path.join(__dirname, "./static")))

app.listen(port)

opn('http://localhost:' + port)

