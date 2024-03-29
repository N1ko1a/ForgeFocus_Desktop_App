import { addDays, addMonths, addWeeks, format, subDays, subMonths, subWeeks } from 'date-fns'
import { useState } from 'react'
import { AiOutlineLeft, AiOutlineRight } from 'react-icons/ai'
import MonthView from './MonthView'
import DayView from './DayView'
import WeekView from './WeekView'

function Kalendar(): JSX.Element {
  const [currentDay, setCurrentDay] = useState(new Date())
  const [buttMonth, setButtMonth] = useState(true)
  const [buttWeek, setButtWeek] = useState(false)
  const [buttDay, setButtDay] = useState(false)

  const handleMonthPrev = () => {
    setCurrentDay(subMonths(currentDay, 1))
  }

  const handleMonthNext = () => {
    setCurrentDay(addMonths(currentDay, 1))
  }
  const handleButtMonth = () => {
    setButtMonth(true)
    setButtWeek(false)
    setButtDay(false)
  }

  const handleButtWeek = () => {
    setButtMonth(false)
    setButtWeek(true)
    setButtDay(false)
  }

  const handleButtDay = () => {
    setButtMonth(false)
    setButtWeek(false)
    setButtDay(true)
  }
  const handleDayPrev = () => {
    setCurrentDay(subDays(currentDay, 1))
  }

  const handleDayNext = () => {
    setCurrentDay(addDays(currentDay, 1))
  }

  const handleWeekPrev = () => {
    setCurrentDay(subWeeks(currentDay, 1))
  }

  const handleWeekNext = () => {
    setCurrentDay(addWeeks(currentDay, 1))
  }
  return (
    <div className="hidden lg:flex flex-col gap-2 justify-between mx-auto p-4  w-4/6 bg-black/40  h-full  rounded-2xl backdrop-blur-sm ">
      <div className=" h-1/6 ">
        <div className="flex   justify-center">
          {buttMonth ? (
            <>
              <button onClick={handleMonthPrev} className="text-center ">
                <AiOutlineLeft className="text-gray-200  text-2xl hover:text-black  transition duration-500 ease-in-out" />
              </button>
              <h1 className="text-center text-gray-200 font-bold text-2xl xl:text-5xl">
                {format(currentDay, 'MMMM yyyy')}
              </h1>
              <button onClick={handleMonthNext} className="text-center text-gray-200  text-2xl">
                <AiOutlineRight className="text-gray-200  text-2xl hover:text-black  transition duration-500 ease-in-out" />
              </button>
            </>
          ) : null}
          {buttDay ? (
            <>
              <button onClick={handleDayPrev} className="text-center ">
                <AiOutlineLeft className="text-gray-200  text-2xl hover:text-black  transition duration-500 ease-in-out" />
              </button>
              <h1 className="text-center text-gray-200 font-bold text-2xl xl:text-5xl">
                {format(currentDay, 'MMMM d yyyy')}
              </h1>
              <button onClick={handleDayNext} className="text-center text-gray-200  text-2xl">
                <AiOutlineRight className="text-gray-200  text-2xl hover:text-black  transition duration-500 ease-in-out" />
              </button>
            </>
          ) : null}
          {buttWeek ? (
            <>
              <button onClick={handleWeekPrev} className="text-center ">
                <AiOutlineLeft className="text-gray-200  text-2xl hover:text-black  transition duration-500 ease-in-out" />
              </button>
              <h1 className="text-center text-gray-200 font-bold text-2xl xl:text-5xl">
                {format(currentDay, 'MMMM yyyy')}
              </h1>
              <button onClick={handleWeekNext} className="text-center text-gray-200  text-2xl">
                <AiOutlineRight className="text-gray-200  text-2xl hover:text-black  transition duration-500 ease-in-out" />
              </button>
            </>
          ) : null}
        </div>
        <div className="flex  mt-4 justify-center text-white   ">
          <button
            className={`flex justify-center items-center mr-2 text-gray-200 p-2 h-6 xl:h-8 rounded-md text-start backdrop-blur-sm hover:bg-black/70 transition duration-500 ease-in-out ${buttMonth ? 'bg-gray-500' : 'bg-gray/30'}`}
            onClick={handleButtMonth}
          >
            Month
          </button>
          <button
            className={`flex justify-center items-center mr-2 text-gray-200 p-2 h-6 xl:h-8 rounded-md text-start backdrop-blur-sm hover:bg-black/70 transition duration-500 ease-in-out ${buttWeek ? 'bg-gray-500' : 'bg-gray/30'}`}
            onClick={handleButtWeek}
          >
            Week
          </button>
          <button
            className={`flex justify-center items-center mr-2 text-gray-200 p-2 h-6 xl:h-8 rounded-md text-start backdrop-blur-sm hover:bg-black/70 transition duration-500 ease-in-out ${buttDay ? 'bg-gray-500' : 'bg-gray/30'}`}
            onClick={handleButtDay}
          >
            Day
          </button>
        </div>
      </div>
      <div className="h-3/4 xl:h-5/6">
        {buttMonth ? <MonthView current={currentDay} /> : null}
        {buttDay ? <DayView current={currentDay} /> : null}
        {buttWeek ? <WeekView current={currentDay} /> : null}
      </div>
    </div>
  )
}
export default Kalendar
