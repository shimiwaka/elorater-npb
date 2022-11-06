import './App.css';
import Vote from './pages/Vote'
import Ranking from './pages/Ranking'
import Player from './pages/Player'
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
        <Route path="player/:id" element={<Player />} />
      </Routes>
    </div>
  );
}

export default App;
