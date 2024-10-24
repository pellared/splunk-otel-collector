import Table from '../components/Table.tsx'
import example from "../exampleData/examle.ts"
function ServiceDiscovery() {
    return ( 
        <>
            <h1>ServiceDiscovery</h1>
            <Table data = {example}></Table>
        </>
     );
}

export default ServiceDiscovery;