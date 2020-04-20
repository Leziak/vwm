import React, { useState } from 'react'
import logo from './logo.svg'
import './App.css'

import SelectDomain from './components/SelectDomain/SelectDomain'


function App() {
  const [domain, setDomain] = useState('House_Targaryen')
  return (
    <>
      <SelectDomain/>
    </>
  )
}

export default App
