const express = require('express')
const router = express.Router()
const Events = require('../model/events')

//get all
router.get('/', async (req, res) => {
  try {
    const events = await Events.find()
    res.status(200).json(events)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//get one
router.get('/:id', getEvent, async (req, res) => {
  try {
    res.status(200).json(res.event)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//creating one
router.post('/', async (req, res) => {
  if (!req.body.Date) {
    return res.status(400).json({ message: 'Morate da unesete index' })
  }
  if (!req.body.Title) {
    return res.status(400).json({ message: 'Morate da unesete naziv' })
  }
  if (!req.body.FromDate) {
    return res.status(400).json({ message: 'Morate da unesete od vreme' })
  }
  if (!req.body.ToDate) {
    return res.status(400).json({ message: 'Morate da unesete do vreme' })
  }

  const event = new Events({
    Date: req.body.Date,
    Title: req.body.Title,
    FromDate: req.body.FromDate,
    ToDate: req.body.ToDate
  })
  try {
    const newEvent = await event.save()
    res.status(200).json({ message: 'Event uspesno dodat' })
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//update one
router.patch('/:id', getEvent, async (req, res) => {
  if (req.body.Date != null) {
    res.event.Date = req.body.Date
  }
  if (req.body.Title != null) {
    res.event.Title = req.body.Title
  }
  try {
    const updateEvent = await res.event.save()
    res.status(200).json(updateEvent)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//delete one
router.delete('/:id', getEvent, async (req, res) => {
  try {
    await res.event.deleteOne()
    res.status(200).json({ message: 'Uspesno izbrisan' })
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

async function getEvent(req, res, next) {
  let event
  try {
    event = await Events.findById(req.params.id)
    if (event == null) {
      return res.status(404).json({ message: 'Cannot find event' })
    }
  } catch (err) {
    return res.status(500).json({ message: err.message })
  }
  res.event = event
  next()
}

module.exports = router
