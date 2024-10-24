import './App.css'
import Table from './components/Table'
import example from "./exampleData/examle.ts"
function App() {

  return (
    <>
    <h1 className="text-3xl font-bold underline">
      Hello world!
    </h1>
      <Table data = {example}></Table>
    </>
  )
}

export default App
