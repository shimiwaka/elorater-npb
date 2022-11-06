import axios from "axios";
import React from 'react';
import VotePlayer from './components/VotePlayer'

const targetURL: string = "http://192.168.0.4:9999/";

const Vote = () => {
  const [token, setToken] = React.useState(null);
  const [player1, setPlayer1] = React.useState(null);
  const [player2, setPlayer2] = React.useState(null);

  React.useEffect(() => {
    axios.get(targetURL + "pick-up").then((response) => {
      if(response.data.error) {
        console.log("Error occured");
      }

      setToken(response.data.token);
      setPlayer1(response.data.player1);
      setPlayer2(response.data.player2);
    });
  }, []);

  if (!player1 || !player2) return ( <div>Loading</div>);

  function select( num : number ){
    axios.get(targetURL + "vote?c=" + num + "&token=" + token).then((response) => {
      if(response.data.error) {
        console.log("Error occured");
      }
      axios.get(targetURL + "pick-up").then((response) => {
        if(response.data.error) {
          console.log("Error occured");
        }
  
        setToken(response.data.token);
        setPlayer1(response.data.player1);
        setPlayer2(response.data.player2);
      });
    });
  }

  return (
    <div>
      <p><VotePlayer player={player1} num={1} onClickFunc={(i : number) => select(i)}/></p>
      <p><VotePlayer player={player2} num={2} onClickFunc={(i : number) => select(i)}/></p>
    </div>
  )
}

export default Vote;
