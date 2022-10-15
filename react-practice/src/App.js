import logo from "./logo.svg";
import "./App.css";
import Form from "./Form.js";
import { useState } from "react";


function App() {
  const [age, setAge] = useState(0);
  const onSubmit = async (name, email, age) => {
    const response = await fetch(
      "https://testuttc-a485d-default-rtdb.asia-southeast1.firebasedatabase.app/tweets.json",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
      const data = await response.json();
      const obj = Object.values(data).filter((v) => v.name === "inada")[0];
      setAge(parseFloat(obj.age) + 10);
  };

  // const onSubmit = async (name, email, age) => {
  //   const response = await fetch(
  //     "https://testuttc-a485d-default-rtdb.asia-southeast1.firebasedatabase.app//tweets.json",
  //     {
  //       method: "POST",
  //       headers: {
  //         "Content-Type": "application/json",
  //       },
  //       body: JSON.stringify({
  //         name,
  //         email,
  //         age,
  //       }),
  //     }
  //   );
  //   console.log("response is...", response);
  // };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Form onSubmit={onSubmit}/>
        <p>
          {age}
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