import React, { useState, useEffect, useRef } from 'react'
import { FaPlay, FaPause } from 'react-icons/fa'
import { VscDebugRestart } from 'react-icons/vsc'
import Timer from './Timer'

const PomodoraTimer = () => {
  const [isActive, setIsActive] = useState(false)
  const [isWorkButtonClicked, setIsWorkButtonClicked] = useState(true)
  const [isRestButtonClicked, setIsRestButtonClicked] = useState(false)
  const [isResetClicked, setIsResetClicked] = useState(false)

  const toggleTimer = () => {
    setIsActive((prevIsActive) => !prevIsActive)
  }

  const handleWorkButton = () => {
    setIsWorkButtonClicked(true)
    setIsRestButtonClicked(false)
  }
  const handleRestButton = () => {
    setIsRestButtonClicked(true)
    setIsWorkButtonClicked(false)
  }
  const handleIsActive = (value) => {
    setIsActive(value)
  }

  const handleIsWorkButtonClicked = (value) => {
    setIsWorkButtonClicked(value)
    setIsRestButtonClicked(!value)
  }
  const handleIsRestButtonClicked = (value) => {
    setIsWorkButtonClicked(!value)
    setIsRestButtonClicked(value)
  }
  const handleResetClicked = () => {
    setIsResetClicked(true)
  }
  const handleIsResetClicked = (value) => {
    setIsResetClicked(value)
  }
  return (
    <div className="flex flex-col gap-2 w-fit p-5 h-5/6 justify-center items-center bg-black/40 rounded-2xl backdrop-blur-sm ">
      {isWorkButtonClicked ? (
        <Timer
          initialHours={'0'}
          initialMinutes={'45'}
          initialSeconds={'0'}
          handleIsActive={handleIsActive}
          isActive={isActive}
          isWorkButtonClicked={isWorkButtonClicked}
          handleIsWorkButtonClicked={handleIsWorkButtonClicked}
          handleIsRestButtonClicked={handleIsRestButtonClicked}
          handleIsResetClicked={handleIsResetClicked}
          isResetClicked={isResetClicked}
        />
      ) : null}
      {isRestButtonClicked ? (
        <Timer
          initialHours={'0'}
          initialMinutes={'15'}
          initialSeconds={'0'}
          handleIsActive={handleIsActive}
          isActive={isActive}
          isWorkButtonClicked={isWorkButtonClicked}
          handleIsWorkButtonClicked={handleIsWorkButtonClicked}
          handleIsRestButtonClicked={handleIsRestButtonClicked}
          handleIsResetClicked={handleIsResetClicked}
          isResetClicked={isResetClicked}
        />
      ) : null}
      <div className="flex w-full h-full justify-between">
        <button
          className={`w-20 md:w-24 lg:w-36 h-10 md: text-sm lg:text-base font-bold mb-5 border-b-2   hover:border-white text-center rounded-xl bg-transparent outline-none    transition duration-500 ease-in-out hover:text-white ${isWorkButtonClicked ? 'text-white border-white' : 'text-gray-400 border-gray-400'}`}
          onClick={handleWorkButton}
        >
          Work
        </button>

        <div className="flex h-full text-gray-300/80 pb-5">
          <button
            onClick={toggleTimer}
            className="text-2xl md:text-3xl lg:text-4xl mr-10 hover:text-white transition duration-500 ease-in-out"
          >
            {isActive ? <FaPause /> : <FaPlay />}
          </button>
          <button
            onClick={handleResetClicked}
            className="text-2xl md:text-3xl lg:text-4xl  hover:text-white transition duration-500 ease-in-out"
          >
            <VscDebugRestart />
          </button>
        </div>
        <button
          className={`w-20 md:w-24 lg:w-36 h-10 md: text-sm lg:text-base font-bold mb-5 border-b-2   hover:border-white text-center rounded-xl bg-transparent outline-none    transition duration-500 ease-in-out hover:text-white ${isRestButtonClicked ? 'text-white border-white' : 'text-gray-400 border-gray-400'}`}
          onClick={handleRestButton}
        >
          Brake
        </button>
      </div>
    </div>
  )
}

export default PomodoraTimer
