import { useState, useEffect, useRef } from 'react'
import { AiOutlineClose } from 'react-icons/ai'
import '../main.css'

function AddButton({ handleCloseButton }) {
  const [nameValue, setNameValue] = useState('')
  const ref = useRef(null)

  const handleClick = () => {
    handleCloseButton(false)
  }
  const handleName = (event) => {
    setNameValue(event.target.value)
  }
  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:3000/buttons', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Name: nameValue
        })
      })
      const data = await response.json()
      if (response.ok) {
        handleCloseButton(false)
        setNameValue('')
      } else {
        console.error('Failed to set event:', data.message)
      }
    } catch (err) {
      console.error('An unexpected error occurred', err)
    }
  }

  // useEffect(() => {
  //   if (date === null) return // No need to focus if editableId is null
  //   if (ref.current) {
  //     ref.current.focus()
  //   }
  // }, [date])

  const handleKeyPress = (event) => {
    if (event.key === 'Enter') {
      handleSubmit()
    }
  }
  return (
    <div
      className=" fixed top-0 left-0  w-full h-full
                flex items-center justify-center z-40 "
    >
      <div className="bg-black/60  w-96 h-fit mb-10  mx-auto   text-black rounded-3xl backdrop-blur-sm ">
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
            placeholder="Add Button Name"
            value={nameValue}
            onChange={handleName}
            onKeyPress={handleKeyPress}
            className="w-4/5 h-10 m-2 bg-transparent rounded-2xl border-b-2   border-gray-500 text-white pl-4 outline-none"
          />
        </div>
        <button
          className=" w-2/5 h-10 mt-5 mb-20 border-b-2  border-gray-500 text-center rounded-xl bg-black/70 outline-none text-gray-400   transition duration-500 ease-in-out hover:text-white "
          onClick={handleSubmit}
        >
          Add Button
        </button>
      </div>
    </div>
  )
}

export default AddButton
