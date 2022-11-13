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

  React.useEffect(() => {
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
  }, []);

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

  const prev = () => {
    if (page <= 0){
      return
    }
    const newPage = page -1;
    setPage(page - 1);
    axios.get(targetURL + "ranking?p=" + newPage).then((response) => {
      if(response.data.error) {
        setError("サーバーエラーが発生しました。しばらくしてから再度お試しください。");
      }

      setPlayers(response.data.players);
    });
  }

  const next = () => {
    const newPage = page + 1;
    setPage(page + 1);
    axios.get(targetURL + "ranking?p=" + newPage).then((response) => {
      if(response.data.error) {
        setError("サーバーエラーが発生しました。しばらくしてから再度お試しください。");
      }

      setPlayers(response.data.players);
    });
  }

  return (
    <div>
      <div className="Navigator">
        <button onClick={() => prev()}> &lt; </button>
        <button onClick={() => next()}> &gt; </button>
      </div>
      <div className="Ranking">
        {players.map(
          (value, i) => 
          {
            return (
              <div className="Ranking-line">
                <div className="Ranking-small-cell" key={i}>
                  {i+page*100+1}位
                </div>
                <div className="Ranking-small-cell" key={i}>
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
