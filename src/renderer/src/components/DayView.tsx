import { eachHourOfInterval, endOfDay, format, isSameDay, startOfDay } from 'date-fns'
import { useState, useEffect } from 'react'

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
  const hourInDay = eachHourOfInterval({
    start: firstHourOfDay,
    end: lastHourOfDay
  })
  const [events, setEvents] = useState<Event[]>([{}])
  return (
    <div className="grid grid-cols-1 gap-2 h-86 mt-6 overflow-auto scrollbar-none">
      {hourInDay.map((hour, index) => {
        return (
          <div
            key={index}
            className="border-2 border-black text-gray-300 p-2 h-28 rounded-md text-strat bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out"
          >
            {format(hour, 'h a')}
            {events
              .filter(
                (event) => isSameDay(event.date, hour) && event.date.getHours() === hour.getHours()
              )
              .map((event) => {
                return <div key={event.title}> {event.title}</div>
              })}
          </div>
        )
      })}
    </div>
  )
}

export default DayView
