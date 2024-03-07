const mongoose = require('mongoose')

const eventsSchema = new mongoose.Schema({
  Date: {
    type: Date,
    require: true
  },
  Title: {
    type: String,
    require: true
  }
})

module.exports = mongoose.model('events', eventsSchema)
