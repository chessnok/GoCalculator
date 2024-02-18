import React, {useEffect, useState} from 'react';
import Expression from "./Expression";
import NotConnectedAlert from "../../../components/NotConnectedAlert";

function ExpressionsList({GetListUrl}) {
    const [isConnected, setIsConnected] = useState(true);
    const fetchData = (data, setData) => {
        fetch(GetListUrl)
            .then(response => {
                if (!response.ok) {
                    setIsConnected(false);
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                setData(data);
                setIsConnected(true);
            })
            .catch(error => {
                console.error('There was an error!', error);
                setIsConnected(false);
            });
    };
    const [data, setData] = useState([]);
    fetchData(data, setData);
    useEffect(() => {
        fetchData(data, setData);
        const interval = setInterval(fetchData, 500);
        return () => clearInterval(interval);
    }, [GetListUrl]);

    return (
        <div>
            {isConnected ? null : (
                <NotConnectedAlert></NotConnectedAlert>
            )}
            <h2>Expressions List</h2>
            {data.sort((a, b) => new Date(b["created_at"]) - new Date(a["created_at"])).map((item, index) => (
                <Expression Expression={item["expression"]} Status={item["status"]} Id={item["id"]}
                            CreatedAt={item["created_at"]} Result={item["result"]}/>
            ))}
        </div>
    );
}

export default ExpressionsList;