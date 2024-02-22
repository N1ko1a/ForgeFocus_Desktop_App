const express = require('express')
const router = express.Router()
const Todo = require('../model/todo')

//get all
router.get('/', async (req, res) => {
  try {
    const todo = await Todo.find()
    res.json(todo)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//get one
router.get('/:id', getTodo, (req, res) => {
  res.json(res.todo)
})

//creating one
router.post('/', async (req, res) => {
  if (!req.body.Content) {
    return res.status(400).json({ message: 'Morate da unesete task' })
  }
  const todo = new Todo({
    Content: req.body.Content
  })
  try {
    const newTodo = await todo.save()
    res.status(200).json({ message: 'Task dodat uspesno' })
  } catch (err) {
    res.status(400).json({ message: err.message })
  }
})

//update one
router.patch('/:id', getTodo, async (req, res) => {
  if (req.body.Content != null) {
    res.todo.Content = req.body.Content
  }
  try {
    const updateTodo = await res.todo.save()
    res.json(updateTodo)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//deleting one
router.delete('/:id', getTodo, async (req, res) => {
  try {
    await res.todo.deleteOne()
    res.status(200).json({ message: 'Uspesno izbrisan' })
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

async function getTodo(req, res, next) {
  let todo
  try {
    todo = await Todo.findById(req.params.id)
    if (todo == null) {
      return res.status(404).json({ message: 'Cannot find todo' })
    }
  } catch (err) {
    return res.status(500).json({ message: err.message })
  }

  res.todo = todo
  next()
}

module.exports = router
