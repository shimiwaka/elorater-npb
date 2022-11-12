import React from 'react';

interface PlayerData {
  name: string;
  birth: string;
  pitchingTotal: any;
  battingTotal: any;
  pitchingCareerHigh: any;
  battingCareerHigh: any;
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

interface BattingStat {
  year: string;
  game: number;
  hit: number;
  avg: number;
  hr: number;
  rbi: number;
  ops: number;
}

type Props = {
  num: number;
  player : PlayerData;
  onClickFunc : (i : number) => void;
}

const ShowPitchingStat = ( {stat} : {stat: PitchingStat} ) => {
  if (stat.game === 0) return null;
  return (
    <div className="Stat-line">
      <div className="Stat-cell">
        {stat.year !== "国内通算" ? stat.year + "年" : "国内通算"}
      </div>
      <div className="Stat-cell">
        {stat.game}試合
      </div>
      <div className="Stat-cell">
        {stat.win}勝
      </div>
      <div className="Stat-cell">
        {stat.lose}敗
      </div>
      <div className="Stat-cell">
       {stat.era.toFixed(2)} 
      </div>
      <div className="Stat-cell">
        {stat.inning}回
      </div>
      <div className="Stat-cell">
        {stat.k}K
      </div>
    </div>
  )
}

const ShowBattingStat = ( {stat} : {stat: BattingStat} ) => {
  if (stat.hit < 50) return null;
  return (
    <div className="Stat-line">
      <div className="Stat-cell">
        {stat.year !== "国内通算" ? stat.year + "年" : "国内通算"}
      </div>
      <div className="Stat-cell">
        {stat.game}試合
      </div>
      <div className="Stat-cell">
        {stat.avg >= 1 ? "1.00" : stat.avg.toFixed(3).slice(1)}
      </div>
      <div className="Stat-cell">
        {stat.hr}本　
      </div>
      <div className="Stat-cell">
        {stat.rbi}打点　
      </div>
      <div className="Stat-cell">
        OPS{stat.ops >= 1 ? stat.ops.toFixed(3) : stat.ops.toFixed(3).slice(1) }
      </div>
      <div className="Stat-cell">
        {stat.hit}安打
      </div>
    </div>
  )
}

const VotePlayer = ( { num, player, onClickFunc } : Props) => {
  return (
    <button onClick={() => onClickFunc(num)}>
      <div>{player.name} ({player.birth} )</div>
      <div className="Stat-table">
        <div>[投手成績]</div>
        <ShowPitchingStat stat={player.pitchingCareerHigh}/>
        <ShowPitchingStat stat={player.pitchingTotal}/>
        <div>[打者成績]</div>
        <ShowBattingStat stat={player.battingCareerHigh}/>
        <ShowBattingStat stat={player.battingTotal}/>
      </div>
    </button>
  )
}

export default VotePlayer;
