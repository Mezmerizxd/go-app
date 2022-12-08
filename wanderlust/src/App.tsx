import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";

export default function App() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    setTimeout(async () => {
      const response = await fetch("http://localhost:3001/api/v1/fetch_count", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => {
        return res.json();
      });
      if (response && !response.error) {
        setCount(response.data.count);
      }
    });
  }, []);

  async function handleCount() {
    const response = await fetch("http://localhost:3001/api/v1/handle_count", {
      method: "POST",
      body: JSON.stringify({ count: count }),
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => {
      return res.json();
    });
    if (response && !response.error) {
      setCount(response.data.count);
    }
  }

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => handleCount()}>Count is {count}</button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </div>
  );
}
