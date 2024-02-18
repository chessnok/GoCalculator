import {Alert} from "react-bootstrap";

function Expression({Id, Expression, Result, CreatedAt, Status}) {
    const variant = {
        "pending": "info", "done": "success", "error": "danger",
    }[Status];
    return (<Alert variant={variant}>
        <Alert.Heading>Expression {Id}</Alert.Heading>
        <p>
            Expression: {Expression}
        </p>
        {Status === "done" && (
            <p>
                Result: {Result}
            </p>
        )}
        <p>
            Created at: {CreatedAt}
        </p>
        <hr/>
    </Alert>);
}


export default Expression;