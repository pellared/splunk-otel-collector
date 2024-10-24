import Table from '../components/Table.tsx'
import example from "../exampleData/examle.ts"
function ServiceSecurity() {
    return ( 
        <>
            <h1>ServiceSecurity</h1>
            <Table data = {example}></Table>
        </>
     );
}

export default ServiceSecurity;