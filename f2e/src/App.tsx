import { useState } from 'react';
import Logo from './assets/logo.png';
import { ManagePage } from './pages';
import './App.css';

function App() {
  const [count, setCount] = useState(0);

  return (
    <div className="App">
      <div>
        <img src={Logo} className="logo" alt="logo" />
        <ManagePage />
      </div>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>count is {count}</button>
      </div>
      <p className="read-the-docs">
        © 2022 <a href="https://github.com/CelloCello">Sero Hsueh</a>
      </p>
    </div>
  );
}

export default App;
