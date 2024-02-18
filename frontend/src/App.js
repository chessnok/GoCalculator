import './App.css';
import Header from "./components/Header";
import {Route, Router, Switch,} from "react-router-dom";
import ConfigPage from "./pages/config";
import {createBrowserHistory} from "history";
import Expressions from "./pages/expressions";
import Agents from "./pages/agents";

function App() {
    var ApiUrl = "http://localhost:80/api";
    return (
        <Router history={createBrowserHistory()}>
            <Header/>
            <Switch>
                <Route path="/config" component={() => <ConfigPage ApiUrl={ApiUrl}></ConfigPage>}/>
                <Route path={"/expressions"} component={() => <Expressions NewExpressionUrl={ApiUrl + "/expression/new"}
                                                                           GetListUrl={ApiUrl + "/expression/list"}></Expressions>}/>
                <Route path={"/agents"} component={() => <Agents GetAgentsUrl={ApiUrl + "/agent/list"}></Agents>}/>
            </Switch>
        </Router>
    );
}

export default App;
