const express = require('express')
const router = express.Router()
const Buttons = require('../model/buttons')

//Get all
router.get('/', async (req, res) => {
  try {
    const button = await Buttons.find()
    res.status(200).json(button)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//get one
router.get('/:id', getButton, async (req, res) => {
  try {
    res.status(200).json(res.button)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//create one
router.post('/', async (req, res) => {
  if (!req.body.Name) {
    return res.status(400).json({ message: 'Morate da unesete ime' })
  }
  const button = new Buttons({
    Name: req.body.Name
  })
  try {
    const newButton = await button.save()
    res.status(200).json({ message: 'Uspesno dodat button' })
  } catch (err) {
    res.status(400).json({ message: err.message })
  }
})

//Update one
router.patch('/:id', getButton, async (req, res) => {
  if (req.body.Name != null) {
    res.button.Name = req.body.Name
  }
  try {
    const updateButton = await res.button.save()
    res.status(200).json({ message: 'Uspeno update button' })
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//deleting one
router.delete('/:id', getButton, async (req, res) => {
  try {
    await res.button.deleteOne()
    res.status(200).json({ message: 'Uspesno izbrisan' })
  } catch (err) {
    res.status(500).json({ message: err.messag })
  }
})

async function getButton(req, res, next) {
  let button
  try {
    button = await Buttons.findById(req.params.id)
    if (button == null) {
      return res.status(404).json({ message: 'Cennot find button' })
    }
  } catch (err) {
    return res.status(500).json({ message: err.message })
  }
  res.button = button
  next()
}

module.exports = router
