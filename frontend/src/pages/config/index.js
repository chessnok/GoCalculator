import React from 'react';
import CalculatorConfigForm from "./components/CalculatorForm";
import AgentConfigForm from "./components/AgentForm";

function ConfigPage({ApiUrl}) {
    return (
        <div>
            <h1>Configuration Page</h1>
            <CalculatorConfigForm getConfigUrl={ApiUrl + "/config/get"}
                                  sendConfigUrl={ApiUrl + "/config/set"}></CalculatorConfigForm><br></br>
            <AgentConfigForm AgentTimeoutCookie={"agent_timeout"}></AgentConfigForm>
        </div>
    );
}

export default ConfigPage;
