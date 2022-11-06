import axios from "axios";
import React from 'react';
import { Routes, Route, Link } from "react-router-dom";

const targetURL: string = "http://192.168.0.4:9999/";

type RankedPlayer = {
  name: string;
  rate: number;
  id: number;
}

const Ranking = () => {
  const [players, setPlayers] = React.useState<RankedPlayer[]>([]);
  const [page, setPage] = React.useState(0);

  React.useEffect(() => {
    axios.get(targetURL + "ranking?p=" + page).then((response) => {
      if(response.data.error) {
        console.log("Error occured");
      }

      setPlayers(response.data.players);
    });
  }, []);

  if(!players) return (<div>loading...</div>);

  const prev = () => {
    if (page <= 0){
      return
    }
    const newPage = page -1;
    setPage(page - 1);
    axios.get(targetURL + "ranking?p=" + newPage).then((response) => {
      if(response.data.error) {
        console.log("Error occured");
      }

      setPlayers(response.data.players);
    });
  }

  const next = () => {
    const newPage = page + 1;
    setPage(page + 1);
    axios.get(targetURL + "ranking?p=" + newPage).then((response) => {
      if(response.data.error) {
        console.log("Error occured");
      }

      setPlayers(response.data.players);
    });
  }

  return (
    <div>
      <div>
        <button onClick={() => prev()}> &lt; </button>
        <button onClick={() => next()}> &gt; </button>
      </div>
      <ul>
        {players.map((value, i) => <li key={i}>{i+page*100+1}位 : {value.rate} : <Link to={`/player/` + value.id}>{value.name}</Link></li>)}
      </ul>
    </div>
  );
}

export default Ranking;
