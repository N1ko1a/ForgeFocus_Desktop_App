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
  const [isLoading, setIsLoading] = useState(true)
  const [isEventSet, setIsEventSet] = useState(false)
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
                const hours = (hour.getHours() + 1).toString().padStart(2, '0')
                const minutes = hour.getMinutes().toString().padStart(2, '0')
                const timestamp = `${hours}:${minutes}`

                return (
                  <div
                    key={index}
                    className="overflow-auto scrollbar-none text-center border-2 border-black text-gray-300 p-2 mb-2 h-20 rounded-md text-strat bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out"
                  >
                    {events
                      .filter(
                        (event) =>
                          isSameDay(event.Date, day) &&
                          timestamp >= event.FromDate &&
                          timestamp <= event.ToDate
                      )
                      .map((event) => {
                        return (
                          <div
                            key={event.Title}
                            className=" mb-1 pl-1 bg-gray-700 rounded-md text-sm truncate"
                          >
                            {event.Title}
                          </div>
                        )
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
