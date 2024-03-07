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
  console.log(isClicked)
  const dayInMonth = eachDayOfInterval({
    start: firstDayofMonth,
    end: lastDayOfMonth
  })
  const [events, setEvents] = useState<Event[]>([{}])
  const startingDayIndex = (getDay(firstDayofMonth) + 6) % 7

  const handleCloseEvent = (value) => {
    setIsClicked(value)
  }

  useEffect(() => {
    setCurrentDay(current)
  }, [current])

  const handleDiv = (index) => {
    setEvents((prevState) => [
      ...prevState,
      { date: setDate(new Date(), index + 1), title: 'test' }
    ])
  }
  const handleClick = (index) => {
    setIsClicked(true)
  }
  return (
    <div className="grid grid-cols-7 gap-2 mt-7   ">
      {isClicked ? <AddEvent handleCloseEvent={handleCloseEvent} /> : null}
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
            className={`border-2 border-black h-28 text-gray-300 bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out rounded-md text-center ${isToday(day) ? 'bg-gray-500' : ''}`}
            onClick={() => handleClick(index)}
          >
            {format(day, 'd')}
            {events
              .filter((event) => isSameDay(event.date, day)) // Update to event.data
              .map((event) => {
                return <div key={event.title}>{event.title}</div>
              })}
          </div>
        )
      })}
    </div>
  )
}
export default MonthView
