import {Alert} from "react-bootstrap";

function Agent({Id, LastPing, Status, ConfigIsUpToDate}) {
    return (
        <Alert variant={Status === 'online' ? ConfigIsUpToDate ? 'success' : 'warning' : 'danger'}>
            <Alert.Heading>Agent {Id}</Alert.Heading>
            <p>
                Last ping: {LastPing}
            </p>
            <hr/>
            <p>
                Status: {Status}
            </p>
            <p>
                Config is up to date: {ConfigIsUpToDate ? 'Yes' : 'No'}
            </p>
        </Alert>
    );

}

export default Agent;