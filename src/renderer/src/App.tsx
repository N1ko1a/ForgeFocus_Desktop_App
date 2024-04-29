import { useEffect } from 'react'
import Kalendar from './components/Kalendar'
import PomodoraTimer from './components/PomodoraTimer'
import ToDo from './components/ToDo'

function App(): JSX.Element {
  const ipcHandle = (): void => window.electron.ipcRenderer.send('ping')

  useEffect(() => {
    window.localStorage.setItem('Count', 0)
  })
  return (
    <div className="bg-[url('/home/neski/Nikola/github/ForgeFocus_Desktop_App/src/renderer/src/assets/pxfuel.jpg')] w-screen h-screen bg-cover bg-center flex flex-col justify-center items-center">
      <div className="flex justify-center items-center text-white w-screen h-1/4">
        <PomodoraTimer />
      </div>
      <div className="flex justify-center w-screen h-3/4 mb-10 mr-0 lg:mr-20">
        <Kalendar />
        <ToDo />
      </div>
    </div>
  )
}

export default App
