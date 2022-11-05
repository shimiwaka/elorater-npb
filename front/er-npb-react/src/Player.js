
import React from 'react';

function PitchingStats(props) {
  if (props.stats.game == 0) return
  return (
    <div>
      {props.stats.year != "国内通算" && props.stats.year + "年" } /
      {props.stats.game} 試合 /
      {props.stats.win} 勝 / 
      {props.stats.lose} 敗 /
      {props.stats.era}
    </div>
  )
}
function BattingStats(props) {
  return (
    <div>
      <div>
        {props.stats.year != "国内通算" && props.stats.year + "年" } /
        {props.stats.game} 試合 / 
        {props.stats.avg} /
        {props.stats.hr}本 /
        {props.stats.rbi}打点 / 
        OPS {props.stats.ops} / 
      </div>
      <div>
        {props.stats.hit}安打
      </div>
    </div>
  )
}


function Player(props) {
  return (
    <button onClick={() => props.onClick(props.num)}>
      <div>{props.player.name}</div>
      <div><PitchingStats stats={props.player.pitchingTotal}/></div>
      <div><PitchingStats stats={props.player.pitchingCareerHigh}/></div>
      <div><BattingStats stats={props.player.battingTotal}/></div>
      <div><BattingStats stats={props.player.battingCareerHigh}/></div>
    </button>
  )
}

export default Player;
