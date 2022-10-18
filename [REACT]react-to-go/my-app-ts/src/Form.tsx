import { useState } from "react";
import { useEffect } from "react";
import axios, { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import user from "./user.json"; //レスポンスのJSON(詳しくは補足で)
const url = "http://localhost:8000";

type USER = typeof user; //画面に表示するユーザー情報の型

type Props = {
  onSubmit: (name: string, age: number) => void;
};

const Form = (props: Props) => {
  const [name, setName] = useState("");
  const [age, setAge] = useState(0);

  useEffect(() => {
    // axios.get<USER>(`${url}/users`)
    //   .then((res) => {
    //     const { name, age } = res;
    //     setName(name);
    //     setAge(age);
    //   })
    //   .catch((e: AxiosError<{ error: string }>) => {
    //     // エラー処理
    //     console.log(e.message);
    //   });
  //   axios.get("http://localhost:8000/users")
      
  //   console.log("aaa")
  // 
  
  axios.get('http://localhost:8000/users')
    .then(response => {
      console.log('status:', response.status);
      console.log('body:', response.data)
    })
  
  // Starting Request: {リクエスト内容} <- consoleに出力される
  
  // Response: {レスポンス内容} <- consoleに出力される

}, []);

  const submit = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    props.onSubmit(name, age);
  };

  return (
    <form className="form">
      <div className="row">
        <label className="label">Name: </label>
        <input
          className="inputBox"
          type={"text"}
          id="name"
          name="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        ></input>
      </div>
      <div className="row">
        <label className="label">Age: </label>
        <input
          className="inputBox"
          type={"number"}
          id="age"
          name="age"
          value={age}
          onChange={(e) => setAge(Number(e.target.value))}
        ></input>
      </div>
      <button className="button" onClick={submit}>
        POST
      </button>
      <div></div>
    </form>
  );
};

export default Form;
