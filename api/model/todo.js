const mongoose = require('mongoose')

const todoSchema = new mongoose.Schema({
  Content: {
    type: String,
    require: true
  },
  Workspace: {
    type: String,
    default: 'Today'
  },
  Compleated: {
    type: Boolean,
    default: false
  }
})

module.exports = mongoose.model('todo', todoSchema)
