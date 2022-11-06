import axios from "axios";
import { useParams } from 'react-router-dom';
import React from 'react';

const targetURL: string = "http://192.168.0.4:9999/";

interface PlayerAllData {
  name: string;
  birth: string;
  bt: string;
  rate: number;
  pitching: PitchingStat[];
  batting: BattingStat[];
}

interface BattingStat {
  year: string;
  game: number;
  hit: number;
  avg: number;
  hr: number;
  rbi: number;
  ops: number;
}

interface PitchingStat {
  year: string;
  game: number;
  win: number;
  lose: number;
  era: number;
  inning: number;
  k: number;
}

const Player = () => {
  const params = useParams();
  const [player, setPlayer] = React.useState<PlayerAllData>();

  React.useEffect(() => {
    axios.get(targetURL + "player/" + params.id).then((response) => {
      setPlayer(response.data);
    });
  }, []);

  if(!player) return (<div>loading...</div>)

  return (
    <ul>
      <li>{player.name}</li>
      <li>{player.birth}</li>
      <li>{player.bt}</li>
      { player.batting.map((value) =>
         <li>{value.year} {value.avg} {value.hr} {value.rbi} {value.ops}</li>)}
      { player.pitching.map((value) =>
          <li>{value.year} {value.win} {value.lose} {value.inning} {value.era}</li>)}
    </ul>
  )
}

export default Player;