import {Button, Form, Toast} from "react-bootstrap";
import {useState} from "react";

function NewExpressionForm({NewExpressionUrl}) {
    const [showToast, setShowToast] = useState(false);
    const [toastMessage, setToastMessage] = useState('');

    function OnSubmit(event) {
        let formData = JSON.stringify({
            expression: event.target.expression.value,
        });
        fetch(NewExpressionUrl, {
            body: formData,
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then((data) => {
                setToastMessage(`Sent an expression, id: ${data["id"]}`);
                setShowToast(true);
            })
            .catch((error) => {
                setToastMessage(`Error, didn't send an expression`);
                setShowToast(true);
            });
        event.preventDefault();
    }

    return (
        <Form onSubmit={OnSubmit}>
            <Form.Group controlId="expression">
                <Form.Label>Expression</Form.Label>
                <Form.Control type="text" name="expression" placeholder="Enter expression"/>
            </Form.Group>
            <Button variant="primary" type="submit">Send Expression</Button>
            <Toast onClose={() => setShowToast(false)} show={showToast} delay={3000} autohide>
                <Toast.Header>
                    <strong className="mr-auto">Response</strong>
                </Toast.Header>
                <Toast.Body>{toastMessage}</Toast.Body>
            </Toast>
        </Form>
    );
}


export default NewExpressionForm;