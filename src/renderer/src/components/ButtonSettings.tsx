import { useState, useEffect, useRef } from 'react'
import { AiOutlineClose, AiFillDelete, AiFillEdit } from 'react-icons/ai'
import { motion, AnimatePresence } from 'framer-motion'
import '../main.css'

function ButtonSettings({
  handleCloseButtonSettings,
  handleButtonSettingsChange,
  handleButtonClick,
  currentButton
}) {
  const [nameValue, setNameValue] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const [editableButtonId, setEditableButtonId] = useState(null)
  const [buttons, setButtons] = useState([])
  const [buttonChange, setButtonChange] = useState(false)
  const ref = useRef(null)

  const handleClick = () => {
    handleCloseButtonSettings(false)
  }
  const handleName = (event) => {
    setNameValue(event.target.value)
  }

  //Uzimam buttons iz baze
  useEffect(() => {
    setIsLoading(true)
    const apiURL = `http://localhost:3000/buttons`
    fetch(apiURL)
      .then((res) => res.json())
      .then((data) => {
        const buttonsResult = data || []
        setButtons(buttonsResult)
        setIsLoading(false)
      })
      .catch((error) => {
        console.log('Error: Ne mogu da uzmem podatke', error)
        setIsLoading(false)
      })

    setButtonChange(false)
  }, [buttonChange])

  //Update button

  const handleKeyDownUpdate = async (event, button) => {
    if (event.key === 'Enter') {
      try {
        const response = await fetch(`http://localhost:3000/buttons/${button._id}`, {
          method: 'PATCH',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            Name: nameValue
          })
        })
        const data = await response.json()

        console.log(data.message)
        if (response.ok) {
          setEditableButtonId(null)
          handleButtonClick(nameValue)
          handleButtonSettingsChange(true)
          setButtonChange(true)
          const response = await fetch(`http://localhost:3000/todo?workspace=${button.Name}`, {
            method: 'PATCH',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              newWorkspace: nameValue
            })
          })
        }
      } catch (error) {
        console.error('An unexpected error occurred', error)
      }
    }
  }

  //Delete
  const handleDelete = async (button) => {
    try {
      const response = await fetch(`http://localhost:3000/buttons/${button._id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
      })
      const data = await response.json()

      if (response.ok) {
        console.log(data.message)

        const anotherResponse = await fetch(`http://localhost:3000/todo?workspace=${button.Name}`, {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({})
        })
        if (currentButton == button.Name) {
          handleButtonClick('Today')
        }
        handleButtonSettingsChange(true)
        setButtonChange(true)
        console.log(buttons.length)
        if (buttons.length == 1) {
          handleCloseButtonSettings(false)
        }
      }
    } catch (error) {
      console.error('An unexpected error occurred', error)
    }
  }

  const handleUpdate = (button) => {
    setEditableButtonId(button._id)
    setNameValue(button.Name)
  }
  return (
    <div
      className=" fixed top-16 left-0  w-full h-1/2
                flex  justify-center z-40 "
    >
      <div className=" flex flex-col items-center bg-black/60  w-full h-full mb-10   mx-auto   text-black rounded-3xl backdrop-blur-sm ">
        <div className="flex justify-end w-11/12 h-5 mt-5">
          <AiOutlineClose
            className=" w-5 h-5 text-gray-500 hover:text-white    transition duration-500 ease-in-out cursor-pointer"
            onClick={() => handleClick()}
          />
        </div>
        <div
          className={`flex flex-col bg-transparent overflow-auto scrollbar-none  w-11/12  h-3/4 mt-8 rounded-lg `}
        >
          <AnimatePresence>
            {buttons.map((button) => (
              <motion.div
                key={button._id}
                className="flex justify-between p-2 w-full h-10 bg-gray/30  backdrop-blur-sm rounded-lg text-white  mt-1 mb-1   "
                initial={{ opacity: 0, scale: 0.5 }}
                animate={{ opacity: 1, scale: 1 }}
                exit={{ opacity: 0, scale: 0.5 }}
                transition={{
                  duration: 0.4,
                  delay: 0.1,
                  ease: [0, 0.71, 0.2, 1.01]
                }}
              >
                <div className="flex w-9/12 justify-start items-start">
                  {editableButtonId == button._id ? (
                    <input
                      ref={ref}
                      type="text"
                      placeholder={button.Name}
                      value={nameValue}
                      onChange={handleName}
                      onKeyDown={(event) => handleKeyDownUpdate(event, button)}
                      onBlur={() => setEditableButtonId(null)}
                      className="w-full h-full bg-transparent rounded-lg text-white pl-4 outline-none"
                    />
                  ) : (
                    <h1 className="text-gray-200 truncate">{button.Name}</h1>
                  )}
                </div>
                <div className="flex w-fit h-full">
                  <button onClick={() => handleUpdate(button)}>
                    <AiFillEdit className="flex justify-center text-gray-400 hover:text-white items-center mr-2  transition duration-500 ease-in-out" />
                  </button>
                  <button onClick={() => handleDelete(button)}>
                    <AiFillDelete className="flex justify-center items-center text-gray-400 hover:text-white transition duration-500 ease-in-out" />
                  </button>
                </div>
              </motion.div>
            ))}
          </AnimatePresence>
        </div>
      </div>
    </div>
  )
}

export default ButtonSettings
