import { useState, useEffect, useRef } from 'react'
import { AiOutlineClose } from 'react-icons/ai'
import { motion, AnimatePresence } from 'framer-motion'
import '../main.css'

function AddEvent({ handleCloseEvent, date, handleEventSet, fromFirstValue, toFirstValue }) {
  const [titleValue, setTitleValue] = useState('')
  const [fromValue, setFromValue] = useState(fromFirstValue)
  const [toValue, setToValue] = useState(toFirstValue)
  const ref = useRef(null)
  const fromRef = useRef(null)
  const toRef = useRef(null)

  const handleClick = () => {
    handleCloseEvent(false)
  }
  const handleTitle = (event) => {
    setTitleValue(event.target.value)
  }
  const handleFrom = (event) => {
    setFromValue(event.target.value)
  }
  const handleTo = (event) => {
    setToValue(event.target.value)
  }
  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:3030/event', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Date: date,
          Title: titleValue,
          FromDate: fromValue,
          ToDate: toValue
        })
      })
      const data = await response.json()
      if (response.ok) {
        handleCloseEvent(false)
        handleEventSet(true)
        setTitleValue('')
        // window.location.reload()
      } else {
        console.error('Failed to set event:', data.message)
      }
    } catch (err) {
      console.error('An unexpected error occurred', err)
    }
  }

  useEffect(() => {
    if (date === null) return // No need to focus if editableId is null
    if (ref.current) {
      ref.current.focus()
    }
  }, [date])

  const handleKeyPress = (event) => {
    if (event.key === 'Enter') {
      if (event.target === ref.current) {
        fromRef.current.focus()
      } else if (event.target === fromRef.current) {
        toRef.current.focus()
      } else if (event.target === toRef.current) {
        handleSubmit()
      }
    }
  }
  return (
    <div
      className=" fixed top-0 left-0  w-full h-full
                flex items-center justify-center z-40 "
    >
      <motion.div
        initial={{ opacity: 0, scale: 0.5 }}
        animate={{ opacity: 1, scale: 1 }}
        transition={{
          duration: 0.4,
          delay: 0.1,
          ease: [0, 0.71, 0.2, 1.01]
        }}
        className="bg-black/60  w-96 h-fit mb-10  mx-auto   text-black rounded-3xl backdrop-blur-sm "
      >
        <div className="w-full h-5 mr-10 relative mb-5 ">
          <AiOutlineClose
            className="w-5 h-5 absolute text-gray-500 hover:text-white top-8 right-5   transition duration-500 ease-in-out cursor-pointer"
            onClick={() => handleClick()}
          />
        </div>
        <div className="flex flex-col justify-center items-center mt-16">
          <input
            ref={ref}
            type="text"
            placeholder="Add Event"
            value={titleValue}
            onChange={handleTitle}
            onKeyPress={handleKeyPress}
            className="w-4/5 h-10 m-2 bg-transparent rounded-2xl border-b-2   border-gray-500 text-white pl-4 outline-none"
          />
          <div className="flex ">
            <input
              ref={fromRef}
              type="time"
              value={fromValue}
              min="1"
              max="12"
              onChange={handleFrom}
              onKeyPress={handleKeyPress}
              className=" text-center w-36 h-10 m-2 bg-transparent rounded-2xl border-b-2   border-gray-500 text-white  outline-none"
            />

            <input
              ref={toRef}
              type="time"
              value={toValue}
              min="1"
              max="12"
              onChange={handleTo}
              onKeyPress={handleKeyPress}
              className=" text-center w-36 h-10 m-2 bg-transparent rounded-2xl border-b-2   border-gray-500 text-white  outline-none"
            />
          </div>
          <button
            className=" w-2/5 h-10 mt-5 mb-20 border-b-2  border-gray-500 text-center rounded-xl bg-black/70 outline-none text-gray-400   transition duration-500 ease-in-out hover:text-white "
            onClick={handleSubmit}
          >
            Add Event
          </button>
        </div>
      </motion.div>
    </div>
  )
}

export default AddEvent
