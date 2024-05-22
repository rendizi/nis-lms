import './App.css';
import Landing from './Pages/Landing/Landing';
import Home from './Pages/Home/Home';
import Task from './Pages/Task/Task';
import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <Routes>
          
          <Route index element={<Landing />} />
          <Route path="/home" element={<Home />} />
          <Route path="/task" element={<Task />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
