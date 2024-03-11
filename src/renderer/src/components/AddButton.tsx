import { useState, useEffect, useRef } from 'react'
import { AiOutlineClose } from 'react-icons/ai'
import '../main.css'

function AddButton({ handleCloseButton, handleButtonChange, focus, handleButtonClick }) {
  const [nameValue, setNameValue] = useState('')
  const ref = useRef(null)

  const handleClick = () => {
    handleCloseButton(false)
  }
  const handleName = (event) => {
    setNameValue(event.target.value)
  }
  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:3000/buttons', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Name: nameValue
        })
      })
      const data = await response.json()
      if (response.ok) {
        handleCloseButton(false)
        handleButtonClick(nameValue)
        setNameValue('')
        handleButtonChange(true)
      } else {
        console.error('Failed to set event:', data.message)
      }
    } catch (err) {
      console.error('An unexpected error occurred', err)
    }
  }

  useEffect(() => {
    if (focus === null) return // No need to focus if editableId is null
    if (ref.current) {
      ref.current.focus()
    }

    // Dodajte slušač događaja 'blur' kada se komponenta mountira
    if (ref.current) {
      ref.current.addEventListener('blur', handleBlur)
    }

    // Uklonite slušač događaja 'blur' kada se komponenta demountira
    return () => {
      if (ref.current) {
        ref.current.removeEventListener('blur', handleBlur)
      }
    }
  }, [focus])

  const handleBlur = () => {
    // Pozovite funkciju za zatvaranje prozora kada input izgubi fokus
    handleCloseButton(false)
  }

  const handleKeyPress = (event) => {
    if (event.key === 'Enter') {
      handleSubmit()
    }
  }
  return (
    <div className="fixed top-[-3rem] right-0 w-fit h-8">
      <input
        ref={ref}
        type="text"
        placeholder="Add Button Name"
        value={nameValue}
        onChange={handleName}
        onKeyPress={handleKeyPress}
        className="w-full h-full   rounded-xl border-b-2  bg-black/60 backdrop-blur-sm text-white pl-4 outline-none"
      />
    </div>
  )
}

export default AddButton
