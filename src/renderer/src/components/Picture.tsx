// import { useState } from 'react'

function Picture(): JSX.Element {
  // const [picture] = useState(window.electron.process.versions)

  return (
    <form
      encType="multipart/form-data"
      action="http://localhost:8080/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  )
}

export default Picture
