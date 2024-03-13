import React, { useState, useEffect, useRef } from 'react'
import { FaPlay, FaPause } from 'react-icons/fa'
import { VscDebugRestart } from 'react-icons/vsc'

const Timer = () => {
  const storedHours = parseInt(window.localStorage.getItem('Hours'))
  const storedMinutes = parseInt(window.localStorage.getItem('Minutes'))
  const storedSeconds = parseInt(window.localStorage.getItem('Seconds'))

  const [hours, setHours] = useState(!isNaN(storedHours) ? storedHours : 0)
  const [minutes, setMinutes] = useState(!isNaN(storedMinutes) ? storedMinutes : 45)
  const [seconds, setSeconds] = useState(!isNaN(storedSeconds) ? storedSeconds : 0)
  const [isActive, setIsActive] = useState(false)
  const [hourIsClicked, setHourIsClicked] = useState(false)
  const [minutesIsClicked, setMinutesIsClicked] = useState(false)
  const [secondsIsClicked, setSecondsIsClicked] = useState(false)
  const [isWorkButtonClicked, setIsWorkButtonClicked] = useState(false)
  const [isRestButtonClicked, setIsRestButtonClicked] = useState(false)
  const hourRef = useRef(null)
  const minutesRef = useRef(null)
  const secondsRef = useRef(null)

  useEffect(() => {
    let intervalId

    if (isActive && (hours > 0 || minutes > 0 || seconds > 0)) {
      intervalId = setInterval(() => {
        if (seconds > 0) {
          setSeconds((prevSeconds) => prevSeconds - 1)
        } else if (minutes > 0) {
          setMinutes((prevMinutes) => prevMinutes - 1)
          setSeconds(59)
        } else if (hours > 0) {
          setHours((prevHours) => prevHours - 1)
          setMinutes(59)
          setSeconds(59)
        }
      }, 1000)
    } else if (hours === 0 && minutes === 0 && seconds === 0) {
      setIsActive(false)
    }

    return () => clearInterval(intervalId)
  }, [isActive, hours, minutes, seconds])

  const toggleTimer = () => {
    setIsActive((prevIsActive) => !prevIsActive)
  }

  const resetTimer = () => {
    setIsActive(false)
    setHours(initialHours)
    setMinutes(initialMinutes)
    setSeconds(initialSeconds)
  }

  const handleHourClick = () => {
    setHourIsClicked(true)
  }

  const handleMinutesClick = () => {
    setMinutesIsClicked(true)
  }

  const handleSecondsClick = () => {
    setSecondsIsClicked(true)
  }

  const handleHoursUpdate = (event) => {
    const newHours = event.target.value
    window.localStorage.setItem('Hours', newHours)
    setHours(newHours)
  }

  const handleMinutesUpdate = (event) => {
    let newMinutes = event.target.value
    newMinutes = newMinutes.slice(0, 2)
    if (newMinutes >= 0 && newMinutes <= 59) {
      window.localStorage.setItem('Minutes', newMinutes)
      setMinutes(newMinutes)
    }
  }

  const handleSecondsUpdate = (event) => {
    let newSeconds = event.target.value
    // Ograničavamo dužinu na dve cifre
    newSeconds = newSeconds.slice(0, 2)
    // Ako je unutar opsega 0-59, čuvamo vrednost i ažuriramo lokalno skladište
    if (newSeconds >= 0 && newSeconds <= 59) {
      window.localStorage.setItem('Seconds', newSeconds)
      setSeconds(newSeconds)
    }
  }

  const handleKeyDownUpdate = (event) => {
    if (event.key === 'Enter') {
      if (hourIsClicked) {
        setHourIsClicked(false)
        setMinutesIsClicked(true)
      } else if (minutesIsClicked) {
        setMinutesIsClicked(false)
        setSecondsIsClicked(true)
      } else if (secondsIsClicked) {
        setSecondsIsClicked(false)
      }
    }
  }

  useEffect(() => {
    if (hourIsClicked) {
      if (hourRef.current) {
        hourRef.current.focus()
      }
    } else if (minutesIsClicked) {
      if (minutesRef.current) {
        minutesRef.current.focus()
      }
    } else if (secondsIsClicked) {
      if (secondsRef.current) {
        secondsRef.current.focus()
      }
    }
  }, [hourIsClicked, minutesIsClicked, secondsIsClicked])

  const handleWorkButton = () => {
    setIsWorkButtonClicked(true)
    setIsRestButtonClicked(false)
  }
  const handleRestButton = () => {
    setIsRestButtonClicked(true)
    setIsWorkButtonClicked(false)
  }

  return (
    <div className="flex flex-col gap-2 w-fit p-5 h-5/6 justify-center items-center bg-black/40 rounded-2xl backdrop-blur-sm ">
      <div className="flex justify-center w-full h-full pt-2 pb-2 text-gray-300/80 font-bold">
        {hourIsClicked ? (
          <input
            ref={hourRef}
            type="number"
            value={hours}
            onChange={handleHoursUpdate}
            onKeyDown={handleKeyDownUpdate}
            onBlur={() => setHourIsClicked(false)}
            className="w-24 h-16 md:w-36 md:h-24 lg:w-44 lg:h-32 bg-transparent text-7xl md:text-8xl lg:text-9xl text-white  outline-none"
          />
        ) : (
          <h1
            onClick={handleHourClick}
            className="w-24 h-16 md:w-36 md:h-24 lg:w-44 lg:h-32  text-7xl md:text-8xl lg:text-9xl"
          >
            {hours < 10 ? '0' + hours : hours}
          </h1>
        )}
        <h1 className="text-7xl md:text-8xl lg:text-9xl">:</h1>
        {minutesIsClicked ? (
          <input
            ref={minutesRef}
            type="number"
            value={minutes}
            onChange={handleMinutesUpdate}
            onKeyDown={handleKeyDownUpdate}
            onBlur={() => setMinutesIsClicked(false)}
            className="w-24 h-16 md:w-36 md:h-24 lg:w-44 lg:h-32  bg-transparent  text-7xl md:text-8xl lg:text-9xl text-white  outline-none"
          />
        ) : (
          <h1
            onClick={handleMinutesClick}
            className="w-24 h-16 md:w-36 md:h-24 lg:w-44 lg:h-32  text-7xl md:text-8xl lg:text-9xl"
          >
            {minutes < 10 && minutes.length < 2 ? '0' + minutes : minutes}
          </h1>
        )}
        <h1 className="text-7xl md:text-8xl lg:text-9xl">:</h1>
        {secondsIsClicked ? (
          <input
            ref={secondsRef}
            type="number"
            value={seconds}
            onChange={handleSecondsUpdate}
            onKeyDown={handleKeyDownUpdate}
            onBlur={() => setSecondsIsClicked(false)}
            className="w-24 h-16 md:w-36 md:h-24 lg:w-44 lg:h-32 bg-transparent text-7xl md:text-8xl lg:text-9xl  text-white  outline-none"
          />
        ) : (
          <h1
            onClick={handleSecondsClick}
            className="w-24 h-16 md:w-36 md:h-24 lg:w-44 lg:h-32  text-7xl md:text-8xl lg:text-9xl"
          >
            {seconds < 10 && seconds.length < 2 ? '0' + seconds : seconds}
          </h1>
        )}
      </div>
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
            onClick={resetTimer}
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

export default Timer
