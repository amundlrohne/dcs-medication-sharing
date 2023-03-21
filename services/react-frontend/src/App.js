import './App.css';

import Login from './pages/Login'
import MedicationForm from './pages/MedicationForm'

import Navbar from './components/Navbar'

import {Route, Routes} from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Navbar />
      <Routes>
        <Route exact path="/login" Component={Login} />
        <Route exact path="/medication-form" Component={MedicationForm} />
      </Routes>
    </div>
  );
}

export default App;
