import { useState, useEffect, useRef } from 'react'
import { AiOutlineClose } from 'react-icons/ai'
import '../main.css'

function EventOptions({
  handleCloseEventOptions,
  date,
  handleEventChange,
  fromFirstValueEvent,
  toFirstValueEvent,
  eventId,
  eventTitle
}) {
  const [titleValue, setTitleValue] = useState(eventTitle)
  const [fromValue, setFromValue] = useState(fromFirstValueEvent)
  const [toValue, setToValue] = useState(toFirstValueEvent)
  const [eventValue, setEventValue] = useState(true)
  const ref = useRef(null)
  const fromRef = useRef(null)
  const toRef = useRef(null)

  const handleClick = () => {
    handleCloseEventOptions(false)
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
  const handleEdit = async () => {
    try {
      const response = await fetch(`http://localhost:3000/event/${eventId}`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Title: titleValue,
          FromDate: fromValue,
          ToDate: toValue
        })
      })
      const data = await response.json()
      if (response.ok) {
        handleCloseEventOptions(false)
        handleEventChange(eventValue)
        setTitleValue('')
        // window.location.reload()
      } else {
        console.error('Failed to update event:', data.message)
      }
    } catch (err) {
      console.error('An unexpected error occurred', err)
    }
  }
  const handleDelete = async () => {
    try {
      const response = await fetch(`http://localhost:3000/event/${eventId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
      })
      const data = await response.json()
      if (response.ok) {
        handleCloseEventOptions(false)
        handleEventChange(eventValue)
      } else {
        console.error('Failed to delete event:', data.message)
      }
    } catch (err) {
      console.error('An unexpected error occurred', err)
    }
  }

  useEffect(() => {
    if (eventId === null) return // No need to focus if editableId is null
    if (ref.current) {
      ref.current.focus()
      // Postavi kursor na kraj teksta u textarea
      ref.current.selectionStart = ref.current.value.length
      ref.current.selectionEnd = ref.current.value.length
    }
  }, [eventId])

  const handleKeyPress = (event) => {
    if (event.key === 'Enter') {
      if (event.target === ref.current) {
        fromRef.current.focus()
      } else if (event.target === fromRef.current) {
        toRef.current.focus()
      } else if (event.target === toRef.current) {
        handleEdit()
      }
    }
  }

  return (
    <div
      className=" fixed top-0 left-0  w-full h-full
                flex items-center justify-center z-40 "
    >
      <div className="bg-black/60  w-2/5 h-fit mb-10  mx-auto   text-black rounded-3xl backdrop-blur-sm ">
        <div className="w-full h-5 mr-10 relative mb-5 ">
          <AiOutlineClose
            className="w-5 h-5 absolute text-gray-500 hover:text-white top-8 right-5   transition duration-500 ease-in-out cursor-pointer"
            onClick={() => handleClick()}
          />
        </div>
        <div className="flex flex-col justify-center w-full h-full items-center mt-16">
          <textarea
            ref={ref}
            value={titleValue}
            onChange={handleTitle}
            onKeyPress={handleKeyPress}
            className={`w-4/5 m-2 bg-transparent rounded-2xl break-words border-b-2 border-gray-500 text-white pl-4 outline-none overflow-auto scrollbar-none transition-height duration-500 ease-in-out ${titleValue.length > 34 ? 'h-20' : 'h-8'}`}
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
          <div className="flex ">
            <button
              className=" w-36 h-10 mt-5 mr-2 mb-20 border-b-2  border-gray-500 text-center rounded-xl bg-black/70 outline-none text-gray-400   transition duration-500 ease-in-out hover:text-white "
              onClick={handleEdit}
            >
              Edit Event
            </button>
            <button
              className=" w-36 h-10 mt-5 mb-20 border-b-2  border-gray-500 text-center rounded-xl bg-black/70 outline-none text-gray-400   transition duration-500 ease-in-out hover:text-white "
              onClick={handleDelete}
            >
              Delete Event
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default EventOptions
