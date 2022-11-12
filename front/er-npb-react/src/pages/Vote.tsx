import axios from "axios";
import React from 'react';
import VotePlayer from './components/VotePlayer'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";

const targetURL: string = process.env.REACT_APP_BASE_URL || "";

const Vote = () => {
  const [token, setToken] = React.useState(null);
  const [player1, setPlayer1] = React.useState(null);
  const [player2, setPlayer2] = React.useState(null);
  const [error, setError] = React.useState<string>("");

  React.useEffect(() => {
    axios.get(targetURL + "pick-up").then((response) => {
      if(response.data.error) {
        setError("サーバーエラーが発生しました。");
        return;
      }

      setToken(response.data.token);
      setPlayer1(response.data.player1);
      setPlayer2(response.data.player2);
    });
  }, []);

  if (error) {
    return (
      <div>
        {error}
      </div>
    )
  }
  if (!player1 || !player2) {
    return (
      <div>
        {error}
        <FontAwesomeIcon icon={faSpinner} />
      </div>
    )
  }

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
