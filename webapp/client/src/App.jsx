import React from 'react';
import logo from './logo.svg';
import './App.scss';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Login, Register } from './components/index';
import Messenger from './components/messenger';


class App extends React.Component {
    render() {
        return (
            <Router>
            <div className="MessengerApp">
                <Switch>
                    <Route exact path="/" component={Login} />
                    <Route path="/signup" component={Register} />
                    <Route path="/messenger" component={Messenger} />
                </Switch>
            </div>
            </Router>
        );
    }
}

export default App