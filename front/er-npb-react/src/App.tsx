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
          どっちがいい選手？
        </Link>
      </span>
      <nav className="Navi">
        <ul className="Menu-group">
          <li className="Menu-item">
           <Link to="/">[投票]</Link>
          </li>
          <li className="Menu-item">
            <Link to="ranking">[ランキング]</Link>
          </li>
          <li className="Menu-item">
            [このサイトは？]
          </li>
        </ul>
      </nav>
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
