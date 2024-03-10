const mongoose = require('mongoose')

const buttonsSchema = new mongoose.Schema({
  Name: {
    type: String,
    require: true
  }
})

module.exports = mongoose.model('buttons', buttonsSchema)
