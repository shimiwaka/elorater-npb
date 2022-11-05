import axios from "axios";
import React from 'react';
import Player from './Player'
import './App.css';

const targetURL = "http://192.168.0.4:9999/";

function App() {
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

  if (!player1 || !player2) return null;

  function select(num){
    axios.get(targetURL + "select?c=" + num + "&token=" + token).then((response) => {
      if(response.data.error) {
        console.log("Error occured");
      }

      setToken(response.data.token);
      setPlayer1(response.data.player1);
      setPlayer2(response.data.player2);
    });
  }

  return (
    <div className="App">
      <p><Player player={player1} num="1" onClick={(i) => select(i)}/></p>
      <p><Player player={player2} num="2" onClick={(i) => select(i)}/></p>
      <p>{token}</p>
    </div>
  );
}

export default App;
