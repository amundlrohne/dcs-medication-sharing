import './App.css';

import Login from './pages/Login'
import Inbox from './pages/Inbox'
import MedicationForm from './pages/MedicationForm'
import ConsentForm from './pages/ConsentForm'

import Navbar from './components/Navbar'

import {Route, Routes} from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Navbar />
      <Routes>
        <Route exact path="/login" Component={Login} />
        <Route exact path="/inbox" Component={Inbox} />
        <Route exact path="/medication-form" Component={MedicationForm} />
        <Route exact path="/consent-form" Component={ConsentForm}/>
      </Routes>
    </div>
  );
}

export default App;
