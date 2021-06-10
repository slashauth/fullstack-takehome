import React from 'react';
import './App.css';
import { useSymbolByIdQuery } from './generated/graphql';

function App() {

  const { data } = useSymbolByIdQuery({
    variables: {
      id: "AAPL",
    },
  });

  return (
    <div className="App">
      <header className="App-header">
        <div>AAPL</div>
        {data?.symbol?.sessions && data.symbol.sessions.length > 0 && data.symbol.sessions.map((sess) => (
          <ul key={sess.time}>
            <li>Time: {sess.time}</li>
            <li>Open: {sess.open / 100}</li>
            <li>Close: {sess.close / 100}</li>
            <li>High: {sess.high / 100}</li>
            <li>Low: {sess.low / 100}</li>
          </ul>
        ))}
      </header>
    </div>
  );
}

export default App;
