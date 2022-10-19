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
