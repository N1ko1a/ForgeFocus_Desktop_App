import { useState } from 'react'
import { AiOutlineClose } from 'react-icons/ai'

function AddEvent({ handleCloseEvent }) {
  const [titleValue, setTitleValue] = useState('')

  const handleClick = () => {
    handleCloseEvent(false)
  }

  const handleTitle = (event) => {
    setTitleValue(event.target.value)
  }
  const handleSubmit = () => {
    window.localStorage.setItem('title', titleValue)
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
            type="text"
            placeholder="Email"
            value={titleValue}
            onChange={handleTitle}
            className="w-4/5 h-10 m-2 bg-transparent rounded-2xl border-b-2   border-gray-500 text-white pl-4 outline-none"
          />
          <button
            className=" w-2/5 h-10 mt-5 mb-20 border-b-2  border-gray-500 text-center rounded-xl bg-black/70 outline-none text-gray-400   transition duration-500 ease-in-out hover:text-white "
            onClick={handleSubmit}
          >
            Add Event
          </button>
        </div>
      </div>
    </div>
  )
}

export default AddEvent
