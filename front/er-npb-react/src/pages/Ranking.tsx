import axios from "axios";
import React from 'react';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";
import { Link } from "react-router-dom";

const targetURL: string = process.env.REACT_APP_BASE_URL || "";

type RankedPlayer = {
  name: string;
  rate: number;
  id: number;
}

const Ranking = () => {
  const [players, setPlayers] = React.useState<RankedPlayer[]>([]);
  const [page, setPage] = React.useState(0);
  const [error, setError] = React.useState<string>("");
  const [query, setQuery] = React.useState<string>("");

  const getRanking = (page : number) => {
    if (query !== "") {
      axios.get(targetURL + "search?p=" + page + "&q=" + query)
      .then((response) => {
        if(response.data.error) {
          setError("サーバーエラーが発生しました。しばらくしてから再度お試しください。");
          return;
        }
        setPlayers(response.data.players);
      })
      .catch((error : any) => {
        setPlayers([]);
      });
    } else {
      axios.get(targetURL + "ranking?p=" + page)
      .then((response) => {
        if(response.data.error) {
          setError("サーバーエラーが発生しました。しばらくしてから再度お試しください。");
          return;
        }
        setPlayers(response.data.players);
      })
      .catch((error : any) => {
        setError("サーバーエラーが発生しました。しばらくしてから再度お試しください。");
      });
    }
  }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  React.useEffect(() => { getRanking(0) }, []);

  if (error) {
    return (
      <div>
        {error}
      </div>
    )
  }

  if (!players) {
    return (
      <div>
        <FontAwesomeIcon icon={faSpinner} />
      </div>
    )
  }

  const setSearchQuery = (query : string) => {
    setQuery(query);
  }

  const prev = () => {
    if (page <= 0){
      return
    }
    const newPage = page -1;
    setPage(page - 1);
    getRanking(newPage);
  }

  const next = () => {
    if (players.length !== 100) {
      return
    }
    const newPage = page + 1;
    setPage(page + 1);
    getRanking(newPage);
  }

  return (
    <div>
      <div className="SearchBox">
        <input placeholder="選手の名前を入力" onChange={(e) => setSearchQuery(e.target.value)}/>
        <button onClick={() => getRanking(0)}>検索</button>
      </div>
      <div className="Navigator">
        <button onClick={() => prev()}> &lt; </button>
        <button onClick={() => next()}> &gt; </button>
      </div>
      <div className="Ranking">
        {players.map(
          (value, i) => 
          {
            return (
              <div className="Ranking-line" key={i}>
                <div className="Ranking-small-cell">
                  {i+page*100+1}位
                </div>
                <div className="Ranking-small-cell">
                  {value.rate}
                </div>
                <div className="Ranking-cell">
                  <Link to={`/player/` + value.id}>{value.name}</Link>
                </div>
              </div> 
            )
          }
        )}
      </div>
      <div className="Navigator">
        <button onClick={() => prev()}> &lt; </button>
        <button onClick={() => next()}> &gt; </button>
      </div>
    </div>
  );
}

export default Ranking;
