import './App.css';
import Landing from './Pages/Landing/Landing';
import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <Routes>
          
          <Route index element={<Landing />} />
          
          <Route path="/" element={<Landing />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
