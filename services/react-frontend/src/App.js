import './App.css';

import Login from './pages/Login'
import Inbox from './pages/Inbox'
import ConsentForm from './pages/ConsentForm'
import MedicationRequestForm from './pages/MedicationRequestForm';
import ViewMedicationRecord from './pages/ViewMedicationRecord';
import NotAuthorized from './pages/NotAuthorized';

import Navbar from './components/Navbar'

import {Route, Routes} from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Navbar />
      <Routes>
        <Route exact path="/login" Component={Login} />
        <Route exact path="/inbox" Component={Inbox} />
        <Route exact path="/consent-form" Component={ConsentForm}/>
        <Route exact path="/medication-form" Component={MedicationRequestForm} />
        <Route exact path="/view-medication-record" Component={ViewMedicationRecord} />
        <Route exact path="/not-authorized" Component={NotAuthorized} />
      </Routes>
    </div>
  );
}

export default App;
