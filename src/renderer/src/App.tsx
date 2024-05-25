import image from "../../../Pictures/wallpaperflare-cropped.jpg"
import Picture from "./components/Picture"

function App(): JSX.Element {

  return (
    <div
      className="w-screen h-screen bg-cover bg-center flex flex-col justify-center items-center"
      style={{ backgroundImage: `url(${image})` }}
    >
      <Picture />
    </div>
  )
}

export default App
