import React from 'react';
import { Routes, Route } from "react-router-dom";
import Home from './pages/Home';
import Login from './pages/Login.jsx';
import Host from './pages/Host.jsx';
import Live from './pages/Live.jsx';

const App = () => {
  return (
    <Routes>
      <Route path='/' element={<Home />} />
      <Route path='/login' element={<Login />} />
      <Route path='/host' element = {<Host />}/>
      <Route path='/live-seminar' element = {<Live />}/>
    </Routes>
  )
}

export default App;