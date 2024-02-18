import React, {useEffect, useState} from 'react';
import {Button, Form} from 'react-bootstrap';
import NotConnectedAlert from "../../../components/NotConnectedAlert";

const CalculatorConfigForm = ({getConfigUrl, sendConfigUrl}) => {
    const [, setConfig] = useState(null);
    const [isEditing, setIsEditing] = useState(false);
    const [isConnected, setIsConnected] = useState(true);
    const [formData, setFormData] = useState({
        add_execution_time: 0,
        sub_execution_time: 0,
        mul_execution_time: 0,
        div_execution_time: 0,
    });
    const fetchConfigData = async () => {
        try {
            const response = await fetch(getConfigUrl);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setConfig(data);
            setIsConnected(true);
            if (!isEditing) {
                setFormData({
                    add_execution_time: data.add_execution_time / 1000000,
                    sub_execution_time: data.sub_execution_time / 1000000,
                    mul_execution_time: data.mul_execution_time / 1000000,
                    div_execution_time: data.div_execution_time / 1000000,
                });
            }
        } catch (error) {
            console.error('Error fetching data:', error);
            setIsConnected(false);
        }
    };

    useEffect(() => {
        if (!isEditing) {
            fetchConfigData();
            const interval = setInterval(fetchConfigData, 3000);
            return () => clearInterval(interval);
        }
    }, [isEditing]);

    const handleInputChange = (e) => {
        const {name, value} = e.target;
        setFormData({...formData, [name]: value});
    };
    const handleSubmit = (e) => {
        const data = {
            add_execution_time: formData.add_execution_time * 1000000,
            sub_execution_time: formData.sub_execution_time * 1000000,
            mul_execution_time: formData.mul_execution_time * 1000000,
            div_execution_time: formData.div_execution_time * 1000000,
        };
        fetch(sendConfigUrl, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                setIsConnected(true);
                setIsEditing(false);
                fetchConfigData();
            })
            .catch((error) => {
                console.error('Error:', error);
                setIsConnected(false);
            });
        e.preventDefault();
    };


    return (
        <Form onSubmit={handleSubmit}>
            {isConnected ? null : (
                <NotConnectedAlert></NotConnectedAlert>
            )}
            <Form.Group controlId="addExecutionTime">
                <Form.Label>Addition Execution Time in seconds</Form.Label>
                <Form.Control
                    type="number"
                    name="add_execution_time"
                    min={0}
                    value={formData.add_execution_time}
                    onChange={handleInputChange}
                    disabled={!isEditing}
                />
            </Form.Group><br></br>
            <Form.Group controlId="subExecutionTime">
                <Form.Label>Subtraction Execution Time in seconds</Form.Label>
                <Form.Control
                    type="number"
                    name="sub_execution_time"
                    min={0}
                    value={formData.sub_execution_time}
                    onChange={handleInputChange}
                    disabled={!isEditing}
                />
            </Form.Group><br></br>
            <Form.Group controlId="mulExecutionTime">
                <Form.Label>Multiplication Execution Time in seconds</Form.Label>
                <Form.Control
                    type="number"
                    name="mul_execution_time"
                    min={0}
                    value={formData.mul_execution_time}
                    onChange={handleInputChange}
                    disabled={!isEditing}
                />
            </Form.Group><br></br>
            <Form.Group controlId="divExecutionTime">
                <Form.Label>Division Execution Time in seconds</Form.Label>
                <Form.Control
                    type="number"
                    name="div_execution_time"
                    min={"0"}
                    value={formData.div_execution_time}
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

export default CalculatorConfigForm;
