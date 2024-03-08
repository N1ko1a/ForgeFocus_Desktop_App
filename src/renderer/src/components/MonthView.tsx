import {
  eachDayOfInterval,
  endOfMonth,
  format,
  getDay,
  isSameDay,
  isToday,
  setDate,
  startOfMonth
} from 'date-fns'
import { useEffect, useState } from 'react'
import AddEvent from './AddEvent'
// import AddEvent from './AddEvent'

interface Event {
  date: Date
  title: string
}

function MonthView({ current }): JSX.Element {
  const [currentDay, setCurrentDay] = useState(current || new Date())
  const WEEKDAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  const firstDayofMonth = startOfMonth(currentDay)
  const lastDayOfMonth = endOfMonth(currentDay)
  const [isClicked, setIsClicked] = useState(false)
  const [isLoading, setIsLoading] = useState(true)
  const [isEventSet, setIsEventSet] = useState(false)
  const [date, setDate] = useState(new Date())
  const dayInMonth = eachDayOfInterval({
    start: firstDayofMonth,
    end: lastDayOfMonth
  })
  const [events, setEvents] = useState<Event[]>([{}])
  const startingDayIndex = (getDay(firstDayofMonth) + 6) % 7
  const [fromFirstValue, setFromFirstValue] = useState('')
  const [toFirstValue, setToFirstValue] = useState('')

  const handleCloseEvent = (value) => {
    setIsClicked(value)
  }
  const handleEventSet = (value) => {
    setIsEventSet(value)
  }

  useEffect(() => {
    setCurrentDay(current)
  }, [current])

  useEffect(() => {
    setIsLoading(true)

    const apiURL = `http://localhost:3000/event`

    fetch(apiURL)
      .then((res) => res.json())
      .then((data) => {
        const eventResult = data || [] // default to an empty array if results is undefine
        // setArtical(articalResults.articles);
        setEvents(eventResult)
        setIsLoading(false)
      })
      .catch((error) => {
        console.log('Error: Ne mogu da uzmem podatke', error)
        setIsLoading(false)
      })
    setIsEventSet(false)
  }, [isEventSet])
  const handleClick = (date) => {
    setIsClicked(true)
    setDate(date)
  }
  useEffect(() => {
    // Function to get the current time in HH:MM format
    const getCurrentTime = () => {
      const now = new Date()
      const hours = now.getHours().toString().padStart(2, '0')
      const minutes = now.getMinutes().toString().padStart(2, '0')
      return `${hours}:${minutes}`
    }
    const getCurrentTimeAndOne = () => {
      const now = new Date()
      // Add one hour to the current hours
      const hours = (now.getHours() + 1).toString().padStart(2, '0')
      const minutes = now.getMinutes().toString().padStart(2, '0')
      return `${hours}:${minutes}`
    }

    // Set the current time as the default value for "From" input field
    setFromFirstValue(getCurrentTime())
    setToFirstValue(getCurrentTimeAndOne())
  }, [])
  return (
    <div className="grid grid-cols-7 gap-2 mt-7   ">
      {isClicked ? (
        <AddEvent
          handleCloseEvent={handleCloseEvent}
          date={date}
          handleEventSet={handleEventSet}
          fromFirstValue={fromFirstValue}
          toFirstValue={toFirstValue}
        />
      ) : null}
      {WEEKDAYS.map((day) => {
        return (
          <div key={day} className="text-center text-gray-300 font-bold">
            {day}
          </div>
        )
      })}
      {Array.from({ length: startingDayIndex }).map((_, index) => {
        return (
          <div
            key={`empty-${index}`}
            className="border-2 border-black h-28 rounded-md bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out  text-center"
          />
        )
      })}
      {dayInMonth.map((day, index) => {
        return (
          <div
            key={index}
            className={`border-2 border-black h-28 text-gray-300 bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out rounded-md text-center overflow-auto scrollbar-none ${isToday(day) ? 'bg-gray-700/70' : ''}`}
            onClick={() => handleClick(day)}
          >
            {format(day, 'd')}
            {events
              .filter((event) => isSameDay(event.Date, day)) // Update to event.data
              .map((event) => {
                return (
                  <div
                    key={event.Title}
                    className="w-full h-fit mb-1 bg-gray-600 rounded-md text-sm truncate"
                  >
                    {event.Title}
                  </div>
                )
              })}
          </div>
        )
      })}
    </div>
  )
}
export default MonthView
