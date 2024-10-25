import ServiceTable from '../components/ServiceTable.tsx'
import exampleVal from "../exampleData/examleVal.ts"
function Example() {
    return ( 
        <>
            <ServiceTable data = {exampleVal}></ServiceTable>
        </>
     );
}

export default Example;