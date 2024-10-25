import { useState, useEffect } from 'react';
import ServiceSecurityType from '../models/ServiceSecurityType.ts';
import emptySec from '../exampleData/emptySec.ts';
import ServiceTable from '../components/ServiceTable.tsx';

function ServiceSecurity() {
    const [data, setData] = useState<ServiceSecurityType[]>([]);
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

                if (Array.isArray(jsonData) && jsonData.length == 0 || jsonData == null) {
                    setData(emptySec);
                }
                else {
                    console.log(jsonData);
                    setData(jsonData);
                }
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
        return <div>{error}</div>;
    }

    return ( 
        <>
            <ServiceTable data={data}></ServiceTable>
        </>
    );
}

export default ServiceSecurity;