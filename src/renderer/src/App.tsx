import Kalendar from './components/Kalendar'
import Timer from './components/Timer'
import ToDo from './components/ToDo'

function App(): JSX.Element {
  const ipcHandle = (): void => window.electron.ipcRenderer.send('ping')

  return (
    <div className="bg-[url('/home/nikola/Nikola/github/Productivity_Desktop_App/src/renderer/src/assets/pxfuel.jpg')] w-screen h-screen bg-cover bg-center flex flex-col justify-center items-center">
      <div className="flex justify-center items-center text-white w-screen h-1/4">
        <Timer />
      </div>
      <div className="flex justify-center w-screen h-3/4 mb-10 mr-0 lg:mr-20">
        <Kalendar />
        <ToDo />
      </div>
    </div>
  )
}

export default App
