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
import EventOptions from './EventOptions'
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
  const [isEventClicked, setIsEventClicked] = useState(false)
  const [eventId, setEventId] = useState(0)
  const [eventTitle, setEventTitle] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const [isEventSet, setIsEventSet] = useState(false)
  const [isEventChange, setIsEventChange] = useState(false)
  const [date, setDate] = useState(new Date())
  const dayInMonth = eachDayOfInterval({
    start: firstDayofMonth,
    end: lastDayOfMonth
  })
  const [events, setEvents] = useState<Event[]>([{}])
  const startingDayIndex = (getDay(firstDayofMonth) + 6) % 7
  const [fromFirstValue, setFromFirstValue] = useState('')
  const [toFirstValue, setToFirstValue] = useState('')
  const [fromFirstValueEvent, setFromFirstValueEvent] = useState('')
  const [toFirstValueEvent, setToFirstValueEvent] = useState('')

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
    setIsEventChange(false)
  }, [isEventSet, isEventChange])

  const handleClick = (date) => {
    setIsClicked(true)
    setDate(date)
  }
  const handleEventClick = (value, value1, value2, value3) => {
    setIsEventClicked(true)
    setFromFirstValueEvent(value)
    setToFirstValueEvent(value1)
    setEventId(value2)
    setEventTitle(value3)
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
            <div className="flex flex-col  items-center">
              {events
                .filter((event) => isSameDay(event.Date, day))
                .map((event) => {
                  // Function to handle event click with arguments
                  const handleEventClickWithArgs = (e) => {
                    e.stopPropagation()
                    handleEventClick(event.FromDate, event.ToDate, event._id, event.Title)
                  }

                  // Function to handle notification click
                  const handleNotificationClick = (Naslov) => {
                    // Calling your test function to show notification
                    Notification.requestPermission().then((perm) => {
                      if (perm === 'granted') {
                        // Ovde možete koristiti ovu notifikaciju ili promeniti po potrebi
                        new Notification(`${Naslov}`, {
                          body: `${event.Title}`
                        })
                      }
                    })
                  }
                  const [hours, minutes] = event.FromDate.split(':')
                  const eventTime = new Date()
                  eventTime.setHours(parseInt(hours))
                  eventTime.setMinutes(parseInt(minutes))

                  // Proveravamo da li je datum događaja danas
                  const currentDate = new Date()
                  const eventDate = new Date(event.Date)
                  const isSameDay =
                    eventDate.getDate() === currentDate.getDate() &&
                    eventDate.getMonth() === currentDate.getMonth() &&
                    eventDate.getFullYear() === currentDate.getFullYear()
                  // Dodajemo 30 minuta na vreme događaja za notifikaciju
                  const notificationTime = new Date(eventTime.getTime() - 30 * 60000)

                  // Trenutno vreme
                  const currentTime = new Date()

                  // Proveravamo vreme za prikaz notifikacije
                  if (isSameDay && notificationTime > currentTime) {
                    setTimeout(() => {
                      const Naslov = 'Event starts in 30 minuts'
                      handleNotificationClick(Naslov)
                    }, notificationTime - currentTime)
                  }

                  // Dodajemo notifikaciju ako je događaj u budućnosti
                  const timeUntilEventStart = eventTime - currentTime
                  if (isSameDay && timeUntilEventStart > 0) {
                    setTimeout(() => {
                      const Naslov = 'Event started'
                      handleNotificationClick(Naslov)
                    }, timeUntilEventStart)
                  }
                  return (
                    <div
                      key={event.Title}
                      className=" w-11/12 h-fit mb-1 bg-gray-700 rounded-md text-sm truncate hover:cursor-pointer"
                      onClick={handleEventClickWithArgs}
                    >
                      {event.Title}
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
export default MonthView
