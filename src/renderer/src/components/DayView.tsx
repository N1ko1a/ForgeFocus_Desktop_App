import { eachHourOfInterval, endOfDay, format, isSameDay, isSameHour, startOfDay } from 'date-fns'
import { useState, useEffect } from 'react'
import AddEvent from './AddEvent'

interface Event {
  date: Date
  title: string
}

function DayView({ current }) {
  const [currentDay, setCurrentDay] = useState(current || new Date())
  useEffect(() => {
    setCurrentDay(current)
  }, [current])
  const firstHourOfDay = startOfDay(currentDay)
  const lastHourOfDay = endOfDay(currentDay)
  const [isLoading, setIsLoading] = useState(true)
  const [isEventSet, setIsEventSet] = useState(false)
  const [isClicked, setIsClicked] = useState(false)
  const [date, setDate] = useState(new Date())
  const hourInDay = eachHourOfInterval({
    start: firstHourOfDay,
    end: lastHourOfDay
  })
  const [events, setEvents] = useState<Event[]>([{}])
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
    console.log(date)
  }
  const handleCloseEvent = (value) => {
    setIsClicked(value)
  }
  const handleEventSet = (value) => {
    setIsEventSet(value)
  }
  return (
    <div className="grid grid-cols-1 gap-2 h-86 mt-6 overflow-auto scrollbar-none">
      {isClicked ? (
        <AddEvent handleCloseEvent={handleCloseEvent} date={date} handleEventSet={handleEventSet} />
      ) : null}
      {hourInDay.map((hour, index) => {
        return (
          <div
            key={index}
            className="border-2 border-black text-gray-300 p-2 h-28 rounded-md text-strat bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out"
            onClick={() => handleClick(hour)}
          >
            {format(hour, 'h a')}
            {events
              .filter((event) => isSameDay(event.Date, hour) && isSameHour(event.Date, hour))
              .map((event) => {
                return <div key={event.Title}> {event.Title}</div>
              })}
          </div>
        )
      })}
    </div>
  )
}

export default DayView
