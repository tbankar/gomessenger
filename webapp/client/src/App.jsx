import React from 'react';
import logo from './logo.svg';
import './App.scss';
import { BrowserRouter as Router, Route, Link, Switch } from 'react-router-dom';
import { Login, Register } from './components/login/index';


class App extends React.Component {
    render() {
        return (
            <Router>
            <div className="MessengerApp">
                <Switch>
                    <Route exact path="/" component={Home} />
                    <Route path="/login" component={Login} />
                    <Route path="/signup" component={Register} />
                </Switch>
            </div>
            </Router>
        );
    }
}

const Home = () => (
    <div>
        <h1>
            Home
        </h1>
    </div>

)



export default App