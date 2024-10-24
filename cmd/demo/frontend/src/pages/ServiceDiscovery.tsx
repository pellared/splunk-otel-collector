import { useEffect, useState } from 'react';
import Table from '../components/Table.tsx';
import Obj from "../models/Obj.ts"

function ServiceDiscovery() {
    const [data, setData] = useState<Obj[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8090/services');
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const jsonData = await response.json();
                console.log(jsonData);
                setData(jsonData);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                } else {
                    setError('An unknown error occurred');
                }
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return ( 
        <>
            <Table data={data}></Table>
        </>
    );
}

export default ServiceDiscovery;
