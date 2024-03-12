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
import AddEvent from './AddEvent'
import EventOptions from './EventOptions'
import { motion, AnimatePresence } from 'framer-motion'

interface Event {
  date: Date
  title: string
}

function WeekView({ current }) {
  const [currentDay, setCurrentDay] = useState(current || new Date())
  const [isLoading, setIsLoading] = useState(true)
  const [isClicked, setIsClicked] = useState(false)
  const [isEventClicked, setIsEventClicked] = useState(false)
  const [fromFirstValue, setFromFirstValue] = useState('')
  const [date, setDate] = useState(new Date())
  const [toFirstValue, setToFirstValue] = useState('')
  const [isEventSet, setIsEventSet] = useState(false)
  const [isEventChange, setIsEventChange] = useState(false)
  const [fromFirstValueEvent, setFromFirstValueEvent] = useState('')
  const [toFirstValueEvent, setToFirstValueEvent] = useState('')
  const [eventId, setEventId] = useState(0)
  const [eventTitle, setEventTitle] = useState('')
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
    setIsEventChange(false)
  }, [isEventSet, isEventChange])

  const handleClick = (date, hour) => {
    setIsClicked(true)
    console.log(date)
    setDate(date)
    const hours = hour.getHours().toString().padStart(2, '0')
    const hoursto = (hour.getHours() + 1).toString().padStart(2, '0')
    const minutes = hour.getMinutes().toString().padStart(2, '0')
    const timestamp = `${hours}:${minutes}`
    const timestampto = `${hoursto}:${minutes}`
    console.log(timestamp)
    console.log(timestampto)
    setFromFirstValue(timestamp)
    setToFirstValue(timestampto)
  }
  const handleCloseEvent = (value) => {
    setIsClicked(value)
  }
  const handleEventSet = (value) => {
    setIsEventSet(value)
  }
  const handleEventChange = (value) => {
    setIsEventChange(value)
  }
  const handleCloseEventOptions = (value) => {
    setIsEventClicked(value)
  }
  const handleEventClick = (value, value1, value2, value3) => {
    setIsEventClicked(true)
    setFromFirstValueEvent(value)
    setToFirstValueEvent(value1)
    setEventId(value2)
    setEventTitle(value3)
  }
  return (
    <div className="grid grid-cols-8 gap-2 h-full  overflow-auto scrollbar-none">
      {isClicked ? (
        <AddEvent
          handleCloseEvent={handleCloseEvent}
          date={date}
          handleEventSet={handleEventSet}
          fromFirstValue={fromFirstValue}
          toFirstValue={toFirstValue}
        />
      ) : null}

      {isEventClicked ? (
        <EventOptions
          handleCloseEventOptions={handleCloseEventOptions}
          date={date}
          handleEventChange={handleEventChange}
          fromFirstValueEvent={fromFirstValueEvent}
          toFirstValueEvent={toFirstValueEvent}
          eventId={eventId}
          eventTitle={eventTitle}
        />
      ) : null}
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
                    onClick={() => handleClick(day, hour)}
                    className="overflow-auto scrollbar-none text-center border-2 border-black text-gray-300 p-2 mb-2 h-20 rounded-md text-strat bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out"
                  >
                    <AnimatePresence>
                      {events
                        .filter(
                          (event) =>
                            isSameDay(event.Date, day) &&
                            timestamp > event.FromDate &&
                            timestamp <= event.ToDate
                        )
                        .map((event) => {
                          const handleEventClickWithArgs = (e) => {
                            e.stopPropagation()
                            handleEventClick(event.FromDate, event.ToDate, event._id, event.Title)
                          }

                          return (
                            <motion.div
                              initial={{ opacity: 0, scale: 0.5 }}
                              animate={{ opacity: 1, scale: 1 }}
                              exit={{ opacity: 0, y: -10 }}
                              transition={{
                                duration: 0.4,
                                delay: 0.1,
                                ease: [0, 0.71, 0.2, 1.01]
                              }}
                              key={event.Title}
                              onClick={handleEventClickWithArgs}
                              className=" hover:cursor-pointer mb-1 pl-1 bg-gray-700 rounded-md text-sm truncate"
                            >
                              {event.Title}
                            </motion.div>
                          )
                        })}
                    </AnimatePresence>
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
