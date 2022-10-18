import logo from "./logo.svg";
import "./App.css";
import Form from "./Form";
import axios from "axios";

function App() {
  const onSubmit = (name: string, age: number) => {
    console.log("onSubmit:", name, " ", age);
    axios.post("http://localhost:8000/user", {
      name: name,
      age: age,
    });
  };

  return (
    <div className="App">
      <header className="App-header">
        <div className="title">User Register</div>
        <Form onSubmit={onSubmit} />
      </header>
    </div>
  );
}

export default App;

// 「何故submit.preventDefault()という処理を入れる必要があるのか」
// 　　　->submitに対するブラウザのデフォルトの反応が効いてしまい、submit毎に入力欄が空になってしまうため

// 「inputに対してvalueで値を指定するとどういうメリットがあるのか」
// 　 ->入力された値を取り出して、useStateに投げることができる
