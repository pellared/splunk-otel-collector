import Table from '../components/Table.tsx'
import exampleVal from "../exampleData/examleVal.ts"
function Example() {
    return ( 
        <>
            <Table data = {exampleVal}></Table>
        </>
     );
}

export default Example;