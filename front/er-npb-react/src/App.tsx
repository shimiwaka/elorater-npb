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
      <div className="Credit">
        <hr/>
        <a href="https://twitter.com/_shimiwaka" target="_blank">開発：しみわか(@_shimiwaka)</a>
        <br/>
        <a href="https://2689web.com/" target="_blank">データ出典：日本プロ野球記録 2689web.com 様</a>
      </div>
    </div>
    </>
  );
}

export default App;
