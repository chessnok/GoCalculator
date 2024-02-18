import React, {useEffect, useState} from "react";
import Agent from "./components/Agent";
import NotConnectedAlert from "../../components/NotConnectedAlert";

function Agents({GetAgentsUrl}) {
    const [isConnected, setIsConnected] = useState(true);
    const [agents, setAgents] = useState([]);
    const fetchAgents = async () => {
        try {
            const response = await fetch(GetAgentsUrl);
            if (!response.ok) {
                setIsConnected(false);
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setAgents(data);
        } catch (error) {
            setIsConnected(false)
            console.error('Error fetching data:', error);
        }
    };
    useEffect(() => {
        fetchAgents();
        const interval = setInterval(fetchAgents, 3000);
        return () => clearInterval(interval);
    }, []);
    return (
        <div>
            <h1>Agents</h1>
            {isConnected ? null : (
                <NotConnectedAlert></NotConnectedAlert>
            )}
            {agents.map((agent) => <Agent ConfigIsUpToDate={agent.config_is_up_to_date} Id={agent.id}
                                          LastPing={agent.last_ping} Status={agent.status}></Agent>)}
        </div>
    );
}

export default Agents;