import {
  eachDayOfInterval,
  isToday,
  endOfWeek,
  format,
  startOfWeek,
  startOfDay,
  endOfDay,
  eachHourOfInterval,
  isSameDay
} from 'date-fns'
import { useState, useEffect } from 'react'

interface Event {
  date: Date
  title: string
}

function WeekView({ current }) {
  const [currentDay, setCurrentDay] = useState(current || new Date())
  useEffect(() => {
    setCurrentDay(current)
  }, [current])

  const [events, setEvents] = useState<Event[]>([{}])
  const weekStart = startOfWeek(currentDay, { weekStartsOn: 1 }) // Set week start to Monday
  const weekEnd = endOfWeek(currentDay, { weekStartsOn: 1 })
  const daysInWeek = eachDayOfInterval({
    start: weekStart,
    end: weekEnd
  })
  const firstHourOfDay = startOfDay(currentDay)
  const lastHourOfDay = endOfDay(currentDay)
  const hourInDay = eachHourOfInterval({
    start: firstHourOfDay,
    end: lastHourOfDay
  })

  return (
    <div className="grid grid-cols-8 gap-2 h-86 mt-6 overflow-auto scrollbar-none">
      <div>
        <div className="border-2 border-black text-gray-300 p-2 h-20 mb-2 rounded-md text-strat bg-gray/30  backdrop-blur-sm ">
          Hours
        </div>

        <div className="grid grid-cols-1 gap-0">
          {hourInDay.map((hour, index) => {
            return (
              <div
                key={index}
                className="border-2 border-black text-gray-300 p-2 h-20 mb-2 rounded-md text-strat bg-gray/30  backdrop-blur-sm "
              >
                {format(hour, 'h a')}
              </div>
            )
          })}
        </div>
      </div>
      {daysInWeek.map((day, index) => {
        return (
          <div className="grid grid-cols-1 gap-2 " key={index}>
            <div
              key={index}
              className={` border-2 border-black text-gray-300 p-2 h-20  rounded-md text-strat bg-gray/30  backdrop-blur-sm  ${isToday(day) ? 'bg-gray-500' : ''}`}
            >
              {format(day, 'E d')}
            </div>
            <div className="grid grid-cols-1 gap-0">
              {hourInDay.map((hour, index) => {
                return (
                  <div
                    key={index}
                    className="border-2 border-black text-gray-300 p-2 mb-2 h-20 rounded-md text-strat bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out"
                  >
                    {events
                      .filter(
                        (event) =>
                          isSameDay(event.date, day) && event.date.getHours() === hour.getHours()
                      )
                      .map((event) => {
                        return <div key={event.title}> {event.title}</div>
                      })}
                  </div>
                )
              })}
            </div>
          </div>
        )
      })}
    </div>
  )
}

export default WeekView
