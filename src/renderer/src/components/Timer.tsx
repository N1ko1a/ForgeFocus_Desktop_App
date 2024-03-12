import React, { useState, useEffect, useRef } from 'react'

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
      setHourIsClicked(false)
      setMinutesIsClicked(false)
      setSecondsIsClicked(false)
    }
  }

  useEffect(() => {
    if (!hourIsClicked) return
    if (hourRef.current) {
      hourRef.current.focus()
    }
  }, [hourIsClicked])

  useEffect(() => {
    if (!minutesIsClicked) return
    if (minutesRef.current) {
      minutesRef.current.focus()
    }
  }, [minutesIsClicked])

  useEffect(() => {
    if (!secondsIsClicked) return
    if (secondsRef.current) {
      secondsRef.current.focus()
    }
  }, [secondsIsClicked])
  return (
    <div>
      <div className="flex">
        {hourIsClicked ? (
          <input
            ref={hourRef}
            type="number"
            placeholder={hours}
            value={hours}
            onChange={handleHoursUpdate}
            onKeyDown={handleKeyDownUpdate}
            onBlur={() => setHourIsClicked(false)}
            className="w-full h-full bg-transparent rounded-lg text-white pl-4 outline-none"
          />
        ) : (
          <h1 onClick={handleHourClick}>{hours < 10 ? '0' + hours : hours}</h1>
        )}
        <h1>:</h1>
        {minutesIsClicked ? (
          <input
            ref={minutesRef}
            type="number"
            placeholder={minutes}
            value={minutes}
            onChange={handleMinutesUpdate}
            onKeyDown={handleKeyDownUpdate}
            onBlur={() => setMinutesIsClicked(false)}
            className="w-full h-full bg-transparent rounded-lg text-white pl-4 outline-none"
          />
        ) : (
          <h1 onClick={handleMinutesClick}>
            {minutes < 10 && minutes.length < 2 ? '0' + minutes : minutes}
          </h1>
        )}
        <h1>:</h1>
        {secondsIsClicked ? (
          <input
            ref={secondsRef}
            type="number"
            placeholder={seconds}
            value={seconds}
            onChange={handleSecondsUpdate}
            onKeyDown={handleKeyDownUpdate}
            onBlur={() => setSecondsIsClicked(false)}
            className="w-full h-full bg-transparent rounded-lg text-white pl-4 outline-none"
          />
        ) : (
          <h1 onClick={handleSecondsClick}>
            {seconds < 10 && seconds.length < 2 ? '0' + seconds : seconds}
          </h1>
        )}
      </div>
      <button onClick={toggleTimer}>{isActive ? 'Pause' : 'Start'}</button>
      <button onClick={resetTimer}>Reset</button>
    </div>
  )
}

export default Timer
