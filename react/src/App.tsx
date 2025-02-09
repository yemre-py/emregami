import "./App.css";
import Button from "./components/Button";

function App() {
  const handleClick = () => {
    console.log("Button clicked");
  };

  return (
    <>
      <Button onClick={handleClick} children="Click me" />
    </>
  );
}

export default App;
