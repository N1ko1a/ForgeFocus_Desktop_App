import React, { useState, useEffect } from 'react'

const Timer = ({ initialHours, initialMinutes, initialSeconds }) => {
  const [hours, setHours] = useState(initialHours || 0)
  const [minutes, setMinutes] = useState(initialMinutes || 0)
  const [seconds, setSeconds] = useState(initialSeconds || 0)
  const [isActive, setIsActive] = useState(false)

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

  const formatTime = (h, m, s) => {
    return `${h < 10 ? '0' + h : h}:${m < 10 ? '0' + m : m}:${s < 10 ? '0' + s : s}`
  }

  return (
    <div>
      <h1>Countdown Timer: {formatTime(hours, minutes, seconds)}</h1>
      <button onClick={toggleTimer}>{isActive ? 'Pause' : 'Start'}</button>
      <button onClick={resetTimer}>Reset</button>
    </div>
  )
}

export default Timer
