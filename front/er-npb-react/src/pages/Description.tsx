import React from 'react';

const Description = () => {
  return (
    <div>
      <div className="Header">
        このサイトは？
      </div>
      <div className="Description">
        <ul>
          <li>ユーザーの投票の結果をもとに、NPBの野球選手のランキングを作るプロジェクトです。</li>
          <li>
            2人の選手をランダムに表示するので、どちらがいい選手と思うかを投票してください。
            その結果によって選手のレートの数値が上下し、ランキングが作られます。
            <ul>
              <li>基本的にレートが近い2人の選手が選出されます。（どうしても対象がいない場合を除く）</li>
            </ul>
          </li>
          <li>あまり名前の知られていない、映像も残っていない昔の選手も、数字を見てそのすごさを知ってほしいと思い作りました。</li>
          <li>ランクづけ方法には「<a href="https://ja.wikipedia.org/wiki/%E3%82%A4%E3%83%AD%E3%83%AC%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0" target="_blank" rel="noreferrer">イロレーティング</a>」を採用しています。</li>
          <li>
            対象となる条件は、打者であれば3000打席以上、投手であれば400試合登板か1500投球回となっています。
            <ul>
              <li>打者の通算打率の基準である4000打数以上、投手の通算防御率の基準である2000投球回をベースに調整しました。</li>
            </ul>
          </li>
          <li>
            選手の成績データは <a href="https://2689web.com/" target="_blank" rel="noreferrer">日本プロ野球記録 2689web.com 様</a>からお借りしました。ありがとうございます！
          </li>
          <li>
            初期値として作者や友人によるテストプレイによる投票、機械的な判断による投票のデータが含まれています。  
          </li>
        </ul>
      </div>
    </div>
  )

}

export default Description;