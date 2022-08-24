import { useState } from "react";
import reactLogo from "./assets/react.svg";
import { ManagePage } from "./pages";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
        <ManagePage />
      </div>
      <h1 className="text-3xl font-bold underline">Hello Tailwind!</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
      <p className="read-the-docs">© 2022 Sero Hsueh</p>
    </div>
  );
}

export default App;
