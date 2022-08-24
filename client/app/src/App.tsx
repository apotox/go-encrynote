import React from 'react';
import './App.css';
import { Routes, Route } from "react-router-dom";
import EncryptPage from './pages/EncryptPage';
import DecryptPage from './pages/DecryptPage';
import { Toast } from './components/Toast';

function App() {
	return (
		<div className="App">
			<Routes>
				<Route path="/" element={<EncryptPage />} />
				<Route path="/note/:id" element={<DecryptPage />} />
			</Routes>

			<Toast />
		</div>
	);
}

export default App;
