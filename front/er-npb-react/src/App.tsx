import './App.css';
import Vote from './pages/Vote'
import Ranking from './pages/Ranking'
import Player from './pages/Player'
import { Routes, Route, Link } from "react-router-dom";

const App = () => {
  return (
    <>
    <header>
      <span className="Title">
        <Link to="/">
          どっちがすごい？
        </Link>
      </span>
      <span className="Menu-group">
        <span className="Menu-item">
          <Link to="/">[投票]</Link>
        </span>
        <span className="Menu-item">
          <Link to="ranking">[ランキング]</Link>
        </span>
        <span className="Menu-item">
          [このサイトは？]
        </span>
      </span>
    </header>
    <div className="App">
      <Routes>
        <Route path="/" element={<Vote />} />
        <Route path="ranking" element={<Ranking />} />
        <Route path="player/:id" element={<Player />} />
      </Routes>
    </div>
    </>
  );
}

export default App;
