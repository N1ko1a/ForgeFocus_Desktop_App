import React, { useState, useEffect, useRef } from 'react'
const Timer = ({
  initialHours,
  initialMinutes,
  initialSeconds,
  isActive,
  isWorkButtonClicked,
  handleIsActive,
  handleIsWorkButtonClicked,
  handleIsRestButtonClicked,
  handleIsResetClicked,
  isResetClicked
}) => {
  console.log(isResetClicked)
  const [hours, setHours] = useState(
    isWorkButtonClicked
      ? window.localStorage.getItem('Work Hours') || initialHours
      : window.localStorage.getItem('Rest Hours') || initialHours
  )
  const [minutes, setMinutes] = useState(
    isWorkButtonClicked
      ? window.localStorage.getItem('Work Minutes') || initialMinutes
      : window.localStorage.getItem('Rest Minutes') || initialMinutes
  )
  const [seconds, setSeconds] = useState(
    isWorkButtonClicked
      ? window.localStorage.getItem('Work Seconds') || initialSeconds
      : window.localStorage.getItem('Rest Seconds') || initialSeconds
  )

  const [hourIsClicked, setHourIsClicked] = useState(false)
  const [minutesIsClicked, setMinutesIsClicked] = useState(false)
  const [secondsIsClicked, setSecondsIsClicked] = useState(false)
  const hourRef = useRef(null)
  const minutesRef = useRef(null)
  const secondsRef = useRef(null)

  useEffect(() => {
    let intervalId
    if (isResetClicked) {
      if (isWorkButtonClicked) {
        setHours(window.localStorage.getItem('Work Hours') || initialHours)
        setMinutes(window.localStorage.getItem('Work Minutes') || initialMinutes)
        setSeconds(window.localStorage.getItem('Work Seconds') || initialSeconds)
        handleIsActive(false)
        handleIsResetClicked(false)
      } else {
        setHours(window.localStorage.getItem('Rest Hours') || initialHours)
        setMinutes(window.localStorage.getItem('Rest Minutes') || initialMinutes)
        setSeconds(window.localStorage.getItem('Rest Seconds') || initialSeconds)
        handleIsActive(false)
        handleIsResetClicked(false)
      }
    }

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
    } else if (hours == 0 && minutes == 0 && seconds == 0) {
      setHours(
        isWorkButtonClicked
          ? window.localStorage.getItem('Work Hours') || initialHours
          : window.localStorage.getItem('Rest Hours') || initialHours
      )
      setMinutes(
        isWorkButtonClicked
          ? window.localStorage.getItem('Work Minutes') || initialMinutes
          : window.localStorage.getItem('Rest Minutes') || initialMinutes
      )
      setSeconds(
        isWorkButtonClicked
          ? window.localStorage.getItem('Work Seconds') || initialSeconds
          : window.localStorage.getItem('Rest Seconds') || initialSeconds
      )
      if (isWorkButtonClicked) {
        let test = window.localStorage.getItem('Count')
        if (test >= 3) {
          window.localStorage.setItem('Count', 0)
          handleIsWorkButtonClicked(true)
          handleIsActive(false)
          Notification.requestPermission().then((perm) => {
            if (perm === 'granted') {
              new Notification('Pomodora Timer', {
                body: 'Full session completed!'
              })
            }
          })
        } else {
          window.localStorage.setItem('Count', parseInt(test) + 1)
          handleIsWorkButtonClicked(false)
          handleIsActive(true)
          Notification.requestPermission().then((perm) => {
            if (perm === 'granted') {
              new Notification('Pomodora Timer', {
                body: `Work session ${window.localStorage.getItem('Count')} compleated!`
              })
            }
          })
        }
      } else {
        handleIsRestButtonClicked(false)
        handleIsActive(true)
        Notification.requestPermission().then((perm) => {
          if (perm === 'granted') {
            new Notification('Pomodora Timer', {
              body: 'Rest session compleated!'
            })
          }
        })
      }
    }

    return () => clearInterval(intervalId)
  }, [isActive, isResetClicked, isWorkButtonClicked, hours, minutes, seconds])

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
    window.localStorage.setItem(isWorkButtonClicked ? 'Work Hours' : 'Rest Hours', newHours)
    if (minutes == 0 && seconds == 0 && newHours != 0) {
      setHours(newHours)
    } else if (minutes != 0 || seconds != 0) {
      setHours(newHours)
    }
  }

  const handleMinutesUpdate = (event) => {
    let newMinutes = event.target.value
    newMinutes = newMinutes.slice(0, 2)
    if (newMinutes >= 0 && newMinutes <= 59) {
      window.localStorage.setItem(isWorkButtonClicked ? 'Work Minutes' : 'Rest Minutes', newMinutes)
      if (hours == 0 && seconds == 0 && newMinutes != 0) {
        setMinutes(newMinutes)
      } else if (hours != 0 || seconds != 0) {
        setMinutes(newMinutes)
      }
    }
  }

  const handleSecondsUpdate = (event) => {
    let newSeconds = event.target.value
    newSeconds = newSeconds.slice(0, 2)
    if (newSeconds >= 0 && newSeconds <= 59) {
      window.localStorage.setItem(isWorkButtonClicked ? 'Work Seconds' : 'Rest Seconds', newSeconds)
      if (hours == 0 && minutes == 0 && newSeconds != 0) {
        setSeconds(newSeconds)
      } else if (hours != 0 || minutes != 0) {
        setSeconds(newSeconds)
      }
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

  return (
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
          {hours < 10 && (hours.length < 2 || hours.length == undefined) ? '0' + hours : hours}
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
          {minutes < 10 && (minutes.length < 2 || minutes.length == undefined) ? '0' + minutes : minutes}
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
          {seconds < 10 && (seconds.length < 2 || seconds.length == undefined) ? '0' + seconds : seconds}
        </h1>
      )}
    </div>
  )
}

export default Timer
