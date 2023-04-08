import React from 'react';

interface PlayerData {
  name: string;
  birth: string;
  pitchingTotal: any;
  battingTotal: any;
  pitchingMLBTotal: any;
  battingMLBTotal: any;
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
  save: number;
  hold: number;
  mlb: boolean;
}

interface BattingStat {
  year: string;
  game: number;
  hit: number;
  avg: number;
  hr: number;
  rbi: number;
  sb: number;
  cs: number;
  obp: number;
  ops: number;
  mlb: boolean;
}

type Props = {
  num: number;
  player : PlayerData;
  onClickFunc : (i : number) => void;
}

const ShowPitchingHeader = ( {player} : {player: PlayerData} ) => {
  if (player.pitchingTotal.game !== 0) {
    return (
      <div className="Stat-header">
        投手成績
      </div>
    )
  } else {
    return (
      <></>
    );
  }
}

const ShowPitchingStat = ( {stat} : {stat: PitchingStat} ) => {
  if (stat.game === 0) return null;
  return (
    <div className="Stat-line">
      <div className="Stat-cell Stat-year">
        {stat.year}<br />
        {stat.mlb && stat.year !== 'MLB通算' ? "(MLB)" : "" }
      </div>
      <div>
        <div className="Stat-line">
          <div className="Stat-cell">
            {stat.game}試合
          </div>
          <div className="Stat-cell">
            {stat.win}勝
          </div>
          <div className="Stat-cell">
            {stat.lose}敗
          </div>
        </div>
        <div className="Stat-line">
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
        <div className="Stat-line">
          <div className="Stat-cell">
            {stat.save}S
          </div>
          <div className="Stat-cell">
            {stat.hold}H
          </div>
          <div className="Stat-cell">
            勝率{(stat.win+stat.lose) === 0 ? "---" : (stat.win > 0 && stat.lose === 0) ? "1.00" : (stat.win/(stat.win+stat.lose)).toFixed(3).slice(1)}
          </div>
        </div>
      </div>
    </div>
  )
}

const ShowBattingHeader = ( {player} : {player: PlayerData} ) => {
  if (player.battingTotal.hit >= 50) {
    return (
      <div className="Stat-header">
        打者成績
      </div>
    )
  } else {
    return (
      <></>
    );
  }
}

const ShowBattingStat = ( {stat} : {stat: BattingStat} ) => {
  if (stat.hit < 50) return null;
  return (
    <div className="Stat-line">
      <div className="Stat-cell Stat-year">
        {stat.year}<br />
        {stat.mlb && stat.year !== 'MLB通算' ? "(MLB)" : "" }
      </div>
      <div>
        <div className="Stat-line">
          <div className="Stat-cell">
            {stat.avg >= 1 ? "1.00" : stat.avg.toFixed(3).slice(1)}
          </div>
          <div className="Stat-cell">
            {stat.hr}本　
          </div>
          <div className="Stat-cell">
            {stat.rbi}打点　
          </div>
        </div>
        <div className="Stat-line">
          <div className="Stat-cell">
            {stat.game}試合
          </div>
          <div className="Stat-cell">
            OBP{stat.obp >= 1 ? "1.00" : stat.obp.toFixed(3).slice(1)}
          </div>
          <div className="Stat-cell">
            OPS{stat.ops >= 1 ? stat.ops.toFixed(3) : stat.ops.toFixed(3).slice(1) }
          </div>
        </div>
        <div className="Stat-line">
          <div className="Stat-cell">
            {stat.hit}安打
          </div>
          <div className="Stat-cell">
            {stat.sb}盗塁
          </div>
          <div className="Stat-cell">
            {stat.cs}盗塁死
          </div>
        </div>
      </div>
    </div>
  )
}

const VotePlayer = ( { num, player, onClickFunc } : Props) => {
  return (
    <button onClick={() => onClickFunc(num)} className="Vote-button">
      <div className="Player-name">{player.name} ({player.birth})</div>
      <div className="Stat-table">
        <ShowPitchingHeader player={player}/>
        <ShowPitchingStat stat={player.pitchingCareerHigh}/>
        <ShowPitchingStat stat={player.pitchingTotal}/>
        <ShowPitchingStat stat={player.pitchingMLBTotal}/>
        <ShowBattingHeader player={player}/>
        <ShowBattingStat stat={player.battingCareerHigh}/>
        <ShowBattingStat stat={player.battingTotal}/>
        <ShowBattingStat stat={player.battingMLBTotal}/>
      </div>
    </button>
  )
}

export default VotePlayer;
