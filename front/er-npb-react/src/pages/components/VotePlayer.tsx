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
    <div>
      <div>
        {stat.year !== "国内通算" ? stat.year + "年" : ""}　
        {stat.game}試合 　
        {stat.win}勝 　
        {stat.lose}敗 　
        {stat.era} 
      </div>
      <div>
        {stat.inning}回　
        {stat.k}奪三振
      </div>
    </div>
  )
}

const ShowBattingStat = ( {stat} : {stat: BattingStat} ) => {
  if (stat.hit < 50) return null;
  return (
    <div>
      <div>
        {stat.year !== "国内通算" ? stat.year + "年" : ""}　
        {stat.avg}　
        {stat.hr}本　
        {stat.rbi}打点　
        OPS{stat.ops}
      </div>
      <div>
        {stat.game}試合　
        {stat.hit}安打
      </div>
    </div>
  )
}

const VotePlayer = ( { num, player, onClickFunc } : Props) => {
  return (
    <button onClick={() => onClickFunc(num)}>
      <div>{player.name}</div>
      <div>{player.birth}</div>
      <div><ShowPitchingStat stat={player.pitchingTotal}/></div>
      <div><ShowPitchingStat stat={player.pitchingCareerHigh}/></div>
      <div><ShowBattingStat stat={player.battingTotal}/></div>
      <div><ShowBattingStat stat={player.battingCareerHigh}/></div>
    </button>
  )
}

export default VotePlayer;
