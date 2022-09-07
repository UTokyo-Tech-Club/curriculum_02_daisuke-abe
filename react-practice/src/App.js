import logo from "./logo.svg";
import "./App.css";
import { useState } from "react";

function App() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [age, setAge] = useState("");

  const onSubmit = (submit) => {
    submit.preventDefault();
    console.log("onSubmit: ", name, age, email);
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <form style={{ display: "flex", flexDirection: "column" }}>
          <label>Name: </label>
          <input
            type={"text"}
            value={name}
            onChange={(e) => setName(e.target.value)}
          ></input>
          <lavel>Age: </lavel>
          <input
            type={"age"}
            value={age}
            onChange={(e) => setAge(e.target.value)}
            ></input>
          <label>Email: </label>
          <input
            type={"email"}
            style={{ marginBottom: 20 }}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          ></input>
          <button onClick={onSubmit}>Submit</button>
        </form>
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;

// 「何故submit.preventDefault()という処理を入れる必要があるのか」
// 　　　->submitに対するブラウザのデフォルトの反応が効いてしまい、submit毎に入力欄が空になってしまうため

// 「inputに対してvalueで値を指定するとどういうメリットがあるのか」　
// 　 ->入力された値を取り出して、useStateに投げることができる