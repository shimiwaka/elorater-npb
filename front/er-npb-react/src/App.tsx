import './App.css';
import Vote from './pages/Vote'
import Ranking from './pages/Ranking'
import { Routes, Route, Link } from "react-router-dom";

const App = () => {
  return (
    <div className="App">
      <div className="navi">
        <Link to="/">Top</Link> | 
        <Link to="ranking">Ranking</Link>
      </div>
      <Routes>
        <Route path="/" element={<Vote />} />
        <Route path="ranking" element={<Ranking />} />
      </Routes>
    </div>
  );
}

export default App;
