require('dotenv').config()
const express = require('express')
const app = express()
const mongoose = require('mongoose')
const cookieParser = require('cookie-parser')
const cors = require('cors')

mongoose.connect(process.env.DATABASE_URL)
const db = mongoose.connection
db.on('error', (error) => console.error(error))
db.once('open', () => {
  console.log('Connected to Database')
})

app.use(express.json())
app.use(cookieParser())

app.use(
  cors({
    origin: ['http://localhost:5173', 'http://localhost:3000'],
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    credentials: true,
    preflightContinue: false,
    optionsSuccessStatus: 204
  })
)

const todoRoute = require('./routes/todo')
app.use('/todo', todoRoute)

const eventRouter = require('./routes/events')
app.use('/event', eventRouter)

const buttonsRouter = require('./routes/buttons')
app.use('/buttons', buttonsRouter)

app.listen(3000, () => console.log('Server strated'))
