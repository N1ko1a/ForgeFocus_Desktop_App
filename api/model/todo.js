const mongoose = require('mongoose')

const todoSchema = new mongoose.Schema({
  Content: {
    type: String,
    require: true
  },
  Compleated: {
    type: Boolean,
    default: false
  }
})

module.exports = mongoose.model('todo', todoSchema)
