import { useEffect, useState } from 'react'
import { AiFillDelete, AiOutlineCheck } from 'react-icons/ai'

function ToDo(): JSX.Element {
  const [tasks, setTasks] = useState([])
  const [inputTask, setInputTask] = useState('')
  const [checkedBox, setCheckedBox] = useState(false)
  const [reload, setReload] = useState(false)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    setIsLoading(true)
    const apiURL = `http://localhost:3000/todo`
    fetch(apiURL)
      .then((res) => res.json())
      .then((data) => {
        const todoResult = data || []
        setTasks(todoResult)
        setIsLoading(false)
      })
      .catch((error) => {
        console.log('Error: Ne mogu da uzmem podatke', error)
        setIsLoading(false)
      })
  }, [tasks])

  const handleInputTask = (event) => {
    setInputTask(event.target.value)
  }
  const handleCheckBoxChange = (e) => {
    setCheckedBox(e.target.checked)
  }

  const handleKeyDown = async (event) => {
    if (event.key === 'Enter') {
      try {
        const response = await fetch('http://localhost:3000/todo', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            Content: inputTask
          })
        })
        const data = await response.json()

        if (response.ok) {
          console.log(data.message)
          setInputTask('')
        }
      } catch (error) {
        console.error('An unexpected error occurred', error)
      }
    }
  }
  const handleDelte = async (index) => {
    try {
      const response = await fetch(`http://localhost:3000/todo/${index}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
      })
      const data = await response.json()

      if (response.ok) {
        console.log(data.message)
      }
    } catch (error) {
      console.error('An unexpected error occurred', error)
    }
  }

  return (
    <div className="bg-black/40 w-1/4 h-2/3 rounded-2xl backdrop-blur-sm flex flex-col justify-between items-center">
      <div className="flex w-11/12 h-10 bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out rounded-lg mt-5">
        <input
          type="text"
          placeholder="Add Task"
          value={inputTask}
          onChange={handleInputTask}
          onKeyDown={handleKeyDown}
          className="w-full h-full  bg-transparent rounded-lg text-white pl-4 outline-none    "
        />
      </div>
      <div className="flex flex-col bg-transparent  w-11/12 h-full mt-5 mb-5 rounded-lg ">
        {tasks.map((task) => (
          <div
            key={task._id}
            className="flex justify-between p-2 w-full h-10 bg-gray/30  backdrop-blur-sm rounded-lg text-white  mt-1 mb-1"
          >
            <div className="flex w-fit  justify-center items-center ">
              <label className="cursor-pointer relative">
                <input
                  type="checkbox"
                  onChange={handleCheckBoxChange}
                  className=" h-5 w-5 border-2 border-gray-200 appearance-none hover:border-gray-600 transition duration-500 ease-in-out  rounded-md mr-2 mt-1"
                />
                <AiOutlineCheck
                  className={`h-5 w-5 text-gray-200 hover:text-gray-600 transition duration-500 ease-in-out absolute left-0 top-1 ${checkedBox ? 'text-opacity-100' : 'text-opacity-0'}`}
                />
              </label>
              <h1 className="text-gray-200">{task.Content}</h1>
            </div>
            <button onClick={() => handleDelte(task._id)}>
              <AiFillDelete className="flex justify-center items-center hover:text-gray-700 transition duration-500 ease-in-out" />
            </button>
          </div>
        ))}
      </div>
    </div>
  )
}
export default ToDo
