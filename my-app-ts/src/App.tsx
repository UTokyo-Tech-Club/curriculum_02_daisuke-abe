import logo from "./logo.svg";
import "./App.css";
import Form from "./Form";

function App() {
  const onSubmit = (name: string, email: string) => {
    console.log("onSubmit:", name, " ", email);
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Form onSubmit={onSubmit}/>
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
