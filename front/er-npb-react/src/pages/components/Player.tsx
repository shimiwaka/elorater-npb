import React from 'react';

interface PlayerData {
  name: string;
  pitchingTotal: any;
  battingTotal: any;
}

type Props = {
  num: number;
  player : PlayerData;
  onClickFunc : (i : number) => void;
}

const PitchingStat = ( {player} : {player: PlayerData} ) => {
  if (player.pitchingTotal.game === 0) return null;
  return (
    <div>
      {player.pitchingTotal.game}試合 /
      {player.pitchingTotal.win}勝 /
      {player.pitchingTotal.lose}敗 /
      {player.pitchingTotal.era} 
    </div>
  )
}

const BattingStat = ( {player} : {player: PlayerData} ) => {
  if (player.battingTotal.hit < 100) return null;
  return (
    <div>
      {player.battingTotal.hit}安打 / 
      {player.battingTotal.avg} / 
      {player.battingTotal.hr}本 /
      {player.battingTotal.rbi}打点 / 
      OPS{player.battingTotal.ops}
    </div>
  )
}


const VotePlayer = ( { num, player, onClickFunc } : Props) => {
  return (
    <button onClick={() => onClickFunc(num)}>
      <div>{player.name}</div>
      <div><PitchingStat player={player}/></div>
      <div><BattingStat player={player}/></div>
    </button>
  )
}

export default VotePlayer;
