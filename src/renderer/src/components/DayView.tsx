import {
  addHours,
  eachHourOfInterval,
  endOfDay,
  format,
  isSameDay,
  isSameHour,
  startOfDay
} from 'date-fns'
import { useState, useEffect } from 'react'
import AddEvent from './AddEvent'
import EventOptions from './EventOptions'
import { motion, AnimatePresence } from 'framer-motion'

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
  const [isEventClicked, setIsEventClicked] = useState(false)
  const [isEventChange, setIsEventChange] = useState(false)
  const [fromFirstValueEvent, setFromFirstValueEvent] = useState('')
  const [toFirstValueEvent, setToFirstValueEvent] = useState('')
  const [eventId, setEventId] = useState(0)
  const [eventTitle, setEventTitle] = useState('')
  const [date, setDate] = useState(new Date())
  const hourInDay = eachHourOfInterval({
    start: firstHourOfDay,
    end: lastHourOfDay
  })
  const [fromFirstValue, setFromFirstValue] = useState('')
  const [toFirstValue, setToFirstValue] = useState('')
  const [events, setEvents] = useState<Event[]>([{}])
  useEffect(() => {
    setIsLoading(true)

    const apiURL = `http://localhost:3030/event`

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

  const handleClick = (date) => {
    setIsClicked(true)
    setDate(date)
    const hours = date.getHours().toString().padStart(2, '0')
    const hoursto = (date.getHours() + 1).toString().padStart(2, '0')
    const minutes = date.getMinutes().toString().padStart(2, '0')
    const timestamp = `${hours}:${minutes}`
    const timestampto = `${hoursto}:${minutes}`
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
    <div className="grid grid-cols-1 gap-2 h-full  overflow-auto scrollbar-none">
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
      {hourInDay.map((hour, index) => {
        const hours = (hour.getHours() + 1).toString().padStart(2, '0')
        const minutes = hour.getMinutes().toString().padStart(2, '0')
        const timestamp = `${hours}:${minutes}`
        return (
          <div
            key={index}
            className="border-2 border-black text-gray-300 p-2 h-20 xl:h-28  rounded-md text-strat bg-gray/30  backdrop-blur-sm hover:bg-black/25  transition duration-500 ease-in-out overflow-auto scrollbar-none"
            onClick={() => handleClick(hour)}
          >
            {format(hour, 'h a')}
            <AnimatePresence>
              {events
                .filter(
                  (event) =>
                    isSameDay(event.Date, hour) &&
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
                      className=" hover:cursor-pointer bg-gray-700  p-1 mb-1 rounded-md  "
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
  )
}

export default DayView
