import './App.css';

import Login from './pages/Login'

import Navbar from './components/Navbar'

import {Route, Routes} from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Navbar />
      <Routes>
        <Route exact path="/login" Component={Login} />
      </Routes>
    </div>
  );
}

export default App;
