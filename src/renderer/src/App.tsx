import Kalendar from './components/Kalendar'
import Timer from './components/Timer'
import ToDo from './components/ToDo'

function App(): JSX.Element {
  const ipcHandle = (): void => window.electron.ipcRenderer.send('ping')

  return (
    <div className="bg-[url('/home/nikola/Nikola/github/Productivity_Desktop_App/src/renderer/src/assets/pxfuel.jpg')] w-screen h-screen bg-cover bg-center flex flex-col justify-center items-center">
      <div className=" text-white w-screen h-1/5">
        <Timer initialHours={1} initialMinutes={30} initialSeconds={10} />
      </div>
      <div className="flex justify-center w-screen h-4/5 mb-10 mr-0 lg:mr-20">
        <Kalendar />
        <ToDo />
      </div>
    </div>
  )
}

export default App
