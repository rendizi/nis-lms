import './App.css';
import Landing from './Pages/Landing/Landing';
import Home from './Pages/Home/Home';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Profile from './Pages/Other/Profile';
import Top from './Pages/Other/Top';
import All from './Pages/Other/All';
import Task from './Pages/Task/Task';

function App() {
  return (
    <BrowserRouter>
      <Routes>
          
          <Route index element={<Landing />} />
          <Route path="/home" element={<Home />} />
          <Route path="/top" element={<Top />} />
          <Route path="/u/*" element={<Profile/>}/>
          <Route path="/t/*" element={<Task/>}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
