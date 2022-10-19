import logo from "./logo.svg";
import "./App.css";
import Form from "./Form";
import axios, { AxiosResponse } from "axios";
import { useState } from "react";
import { useEffect } from "react";

// type json = {
//   id: number,
//   name: string
//   age: number
// };
type hoge = {
  id: string;
  name: string;
  age: number;
}

function App() {
  const onSubmit = (name: string, age: number) => {
    console.log("onSubmit:", name, " ", age);
    axios.post("http://localhost:8000/user", {
      name: name,
      age: age,
    });
  };
  
  const [posts, setPosts] = useState<hoge[]>([]);

  useEffect(() => {
    fetch("http://localhost:8000/users")
      .then((response) => response.json())
      .then((data) => {
        console.log("hoge", data);
        setPosts(data);
      });
  }, []);

  // useEffect(() => {
  //   axios.get<json>("http://localhost:8000/users").then((response) => {
  //       const arr = response.data
  //     });
  // }, []);

  // const arr = ["りんご", "みかん", "ぶどう"];
  return (
    <div className="App">
      <header className="App-header">
        <div className="title">User Register</div>
        <Form onSubmit={onSubmit} />
        <ul>
          {posts.map((user, i) => (
            <li className="list" key={i}>
              {user.name}, {user.age}
            </li>
          ))}
        </ul>
      </header>
    </div>
  );
}

export default App;

// 「何故submit.preventDefault()という処理を入れる必要があるのか」
// 　　　->submitに対するブラウザのデフォルトの反応が効いてしまい、submit毎に入力欄が空になってしまうため

// 「inputに対してvalueで値を指定するとどういうメリットがあるのか」
// 　 ->入力された値を取り出して、useStateに投げることができる
