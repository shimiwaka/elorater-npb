import axios from "axios";
import { useParams } from 'react-router-dom';
import React from 'react';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";

const targetURL: string = process.env.REACT_APP_BASE_URL || "";

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
  mlb: boolean;
}

interface PitchingStat {
  year: string;
  game: number;
  win: number;
  lose: number;
  era: number;
  inning: number;
  k: number;
  mlb: boolean;
}

const isPitcher = (player : PlayerAllData) : boolean => {
  for (let i = 0; i < player.pitching.length; i++){
    if(player.pitching[i].game > 0){
      return true;
    }
  }
  return false;
}

const isBatter = (player : PlayerAllData) : boolean => {
  for (let i = 0; i < player.batting.length; i++){
    if(player.batting[i].hit > 0){
      return true;
    }
  }
  return false;
}

const Player = () => {
  const params = useParams();
  const [player, setPlayer] = React.useState<PlayerAllData>();
  const [error, setError] = React.useState<string>("");

  React.useEffect(() => {
    axios.get(targetURL + "player/" + params.id)
    .then((response) => {
      setPlayer(response.data);
    })
    .catch((error : any) => {
      setError("サーバーエラーが発生しました。しばらくしてから再度お試しください。");
    });
  }, [params.id]);

  if (error) {
    return (
      <div>
        {error}
      </div>
    )
  }

  if (!player) {
    return (
      <div>
        <FontAwesomeIcon icon={faSpinner} />
      </div>
    )
  }
  let i = 0;

  return (
    <div className="Player">
      <div key="name" className="Header">{player.name}</div>
      <div key="birth_bt">{player.birth} {player.bt}</div>

      { isPitcher(player) ? 
        <>
          <div className="Player-header" key="p-header">投手成績</div>
          <div className="Player-line" key={i++}>
            <div className="Player-cell">
              年
            </div>
            <div className="Player-cell">
              勝
            </div>
            <div className="Player-cell">
              敗
            </div>
            <div className="Player-cell">
              防御率
            </div>
            <div className="Player-cell">
              投球回
            </div>
            <div className="Player-cell">
              奪三振
            </div>
          </div>
        </>
         : "" 
      }
        { player.pitching.map((stat) => {
          if(stat.game === 0){
            return null;
          }
          return (
            <div className="Player-line" key={i++}>
              <div className="Player-cell">
                {stat.year}
                {stat.mlb && stat.year !== 'MLB通算' ? "\n(MLB)" : "" }
              </div>
              <div className="Player-cell">
                {stat.win}
              </div>
              <div className="Player-cell">
                {stat.lose}
              </div>
              <div className="Player-cell">
                {stat.era.toFixed(2)}
              </div>
              <div className="Player-cell">
               {stat.inning}
              </div>
              <div className="Player-cell">
               {stat.k}
              </div>
            </div>
          )
       })}

      { isBatter(player) ? 
        <>
          <div className="Player-header" key="b-header">打者成績</div>
          <div className="Player-line" key={i++}>
            <div className="Player-cell">
              年
            </div>
            <div className="Player-cell">
              打率
            </div>
            <div className="Player-cell">
              HR
            </div>
            <div className="Player-cell">
              打点
            </div>
            <div className="Player-cell">
              安打
            </div>
            <div className="Player-cell">
              OPS
            </div>
          </div>
        </>
         : "" 
      }
      { player.batting.map((stat) => {
        if(stat.hit === 0){
          return null;
        }
        return (
          <div className="Player-line" key={i++}>
            <div className="Player-cell">
              {stat.year}
              {stat.mlb && stat.year !== 'MLB通算' ? "\n(MLB)" : "" }
            </div>
            <div className="Player-cell">
              {stat.avg >= 1 ? "1.00" : stat.avg.toFixed(3).slice(1)}
            </div>
            <div className="Player-cell">
              {stat.hr}
            </div>
            <div className="Player-cell">
              {stat.rbi}
            </div>
            <div className="Player-cell">
              {stat.hit}
            </div>
            <div className="Player-cell">
              {stat.ops >= 1 ? stat.ops.toFixed(3) : stat.ops.toFixed(3).slice(1) }
            </div>
          </div>
        )
      })}
    </div>
  )
}

export default Player;