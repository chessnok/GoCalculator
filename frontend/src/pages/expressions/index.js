import NewExpressionForm from "./components/NewExpressionForm";
import ExpressionsList from "./components/ExpressionsList";

function Expressions({GetListUrl, NewExpressionUrl}) {
    return (
        <div>
            <h1>Expressions</h1>
            <NewExpressionForm NewExpressionUrl={NewExpressionUrl}></NewExpressionForm>
            <ExpressionsList GetListUrl={GetListUrl}></ExpressionsList>
        </div>
    );
}

export default Expressions;