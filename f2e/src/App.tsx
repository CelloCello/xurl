import Logo from './assets/logo.png';
import { ManagePage } from './pages';
import './App.css';

function App() {
  return (
    <div className="App">
      <div className="mb-5">
        <img src={Logo} className="logo" alt="logo" />
        <ManagePage />
      </div>
      <p className="read-the-docs">
        © 2022 <a href="https://github.com/CelloCello">Sero Hsueh</a>
      </p>
    </div>
  );
}

export default App;
