import { useState } from "react";

type Props = {
  onSubmit: (name: string, email: string) => void;
}

const Form = (props: Props) => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");

  const submit = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    props.onSubmit(name, email);
  };

  return (
    <form style={{ display: "flex", flexDirection: "column" }}>
      <label>Name: </label>
      <input
        type={"text"}
        value={name}
        onChange={(e) => setName(e.target.value)}
      ></input>
      <label>Email: </label>
      <input
        type={"email"}
        style={{ marginBottom: 20 }}
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      ></input>
      <button onClick={submit}>Submit</button>
    </form>
  );
};

export default Form;