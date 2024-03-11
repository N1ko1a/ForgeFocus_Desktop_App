const express = require('express')
const router = express.Router()
const Todo = require('../model/todo')

router.get('/', async (req, res) => {
  try {
    let query = {} // Inicijalizujemo prazan query objekat

    // Provjeravamo da li je proslijeđen query parametar za pretragu po "Workspace"
    if (req.query.workspace) {
      // Ako jeste, kreiramo regex objekat za pretragu po "Workspace" polju
      const workspaceRegex = new RegExp(req.query.workspace, 'i') // 'i' znači da je pretraga case-insensitive
      query = { Workspace: workspaceRegex } // Postavljamo query da pretražuje "Workspace" polje
    }

    // Izvršavamo upit na bazi podataka sa kreiranim query objektom
    const todo = await Todo.find(query)
    res.status(200).json(todo)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//get one
router.get('/:id', getTodo, async (req, res) => {
  res.json(res.todo)
  try {
    res.status(200).json(res.todo)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//creating one
router.post('/', async (req, res) => {
  if (!req.body.Content) {
    return res.status(400).json({ message: 'Morate da unesete task' })
  }
  if (!req.body.Workspace) {
    return res.status(400).json({ message: 'Morate da unesete Workspace' })
  }
  const todo = new Todo({
    Content: req.body.Content,
    Workspace: req.body.Workspace
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
  if (req.body.Compleated != null) {
    res.todo.Compleated = req.body.Compleated
  }
  if (req.body.Workspace != null) {
    res.todo.Workspace = req.body.Workspace
  }
  try {
    const updateTodo = await res.todo.save()
    res.status(200).json(updateTodo)
  } catch (err) {
    res.status(500).json({ message: err.message })
  }
})

//update many
router.patch('/', async (req, res) => {
  let query = {} // Initialize an empty query object

  // Provjeravamo da li je proslijeđen query parametar za pretragu po "Workspace"
  if (req.query.workspace) {
    // Ako jeste, kreiramo regex objekat za pretragu po "Workspace" polju
    const workspaceRegex = new RegExp(req.query.workspace, 'i') // 'i' znači da je pretraga case-insensitive
    query = { Workspace: workspaceRegex } // Postavljamo query da pretražuje "Workspace" polje
  } else {
    return res.status(400).json({ message: 'Query parameter "workspace" is required.' })
  }

  // Provjeravamo da li je poslana nova vrijednost za "Workspace"
  if (!req.body.newWorkspace) {
    return res
      .status(400)
      .json({ message: 'New value for "Workspace" is required in request body.' })
  }

  try {
    // Update all documents that match the query
    const updateResult = await Todo.updateMany(query, {
      $set: { Workspace: req.body.newWorkspace }
    })

    if (updateResult.modifiedCount > 0) {
      res
        .status(200)
        .json({ message: 'Successfully updated ' + updateResult.nModified + ' documents.' })
    } else {
      res.status(404).json({ message: 'No documents found matching the criteria.' })
    }
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

//delete all
router.delete('/', async (req, res) => {
  try {
    let query = {} // Inicijalizujemo prazan query objekat

    // Provjeravamo da li je proslijeđen query parametar za pretragu po "Workspace"
    if (req.query.workspace) {
      // Ako jeste, kreiramo regex objekat za pretragu po "Workspace" polju
      const workspaceRegex = new RegExp(req.query.workspace, 'i') // 'i' znači da je pretraga case-insensitive
      query = { Workspace: workspaceRegex } // Postavljamo query da pretražuje "Workspace" polje
    }

    // Izvršavamo upit na bazi podataka sa kreiranim query objektom
    const todo = await Todo.deleteMany(query)
      .then(function () {
        // Success
        console.log('Data deleted')
      })
      .catch(function (error) {
        // Failure
        console.log(error)
      })
    res.status(200).json({ message: 'Uspesno svi taskovi izbrisani' })
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
