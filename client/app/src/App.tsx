import React from 'react';
import './App.css';
import { Routes, Route } from "react-router-dom";
import Home from './pages/Home';
import NotePage from './pages/NotePage';
import { Toast } from './components/Toast';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/note/:id" element={<NotePage />} />
      </Routes>

      <Toast />
    </div>
  );
}

export default App;
