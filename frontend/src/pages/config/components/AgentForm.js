import React, {useState} from 'react';
import {Button, Form} from 'react-bootstrap';
import {useCookies} from "react-cookie";

const AgentConfigForm = ({AgentTimeoutCookie}) => {
    const [cookies, setCookie, removeCookie] = useCookies(['cookie-name']);
    const [isEditing, setIsEditing] = useState(false);
    const [formData, setFormData] = useState({
        agent_timeout: cookies[AgentTimeoutCookie] || 0,
    });
    const handleInputChange = (e) => {
        const {name, value} = e.target;
        setFormData({...formData, [name]: value});
    };
    const handleSubmit = (e) => {
        setCookie(AgentTimeoutCookie, formData.agent_timeout, {path: '/'});
        setIsEditing(false);
        e.preventDefault();
    };

    return (
        <Form onSubmit={handleSubmit}>
            <Form.Group controlId="divExecutionTime">
                <Form.Label>Don't show agent after in minutes:</Form.Label>
                <Form.Control
                    type="number"
                    name="agent_timeout"
                    min={1}
                    value={formData.agent_timeout}
                    onChange={handleInputChange}
                    disabled={!isEditing}
                />
            </Form.Group><br></br>
            {isEditing ? (
                <>
                    <Button variant="primary" type="submit">
                        Save
                    </Button>
                    <Button variant="secondary" onClick={() => setIsEditing(false)}>
                        Cancel
                    </Button>
                </>
            ) : (
                <Button variant="primary" onClick={() => setIsEditing(true)}>
                    Edit
                </Button>
            )}
        </Form>
    );
};

export default AgentConfigForm;
