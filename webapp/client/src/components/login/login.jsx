import React from 'react';
import LogoImage from '../../login.svg';
import { Redirect, Link } from 'react-router-dom'
import axios from 'axios'

export class Login extends React.Component {

    constructor(props) {
        super(props);

        this.state ={
            username:null,
            password:null,
            loggedin:false,
            unauth:false,
            sourceipaddr:"",
        };
        this.doLogin = this.doLogin.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    handleChange = (event) => {
        const {name,value} = event.target
        this.setState({[name]:value});
    }

    handleResponseError(response) {
        throw new Error("HTTP error, status = " + response.status);
    }

    handleError(error) {
        console.log(error.message);
    }

    

    doLogin = (event) => {
        event.preventDefault()

        const isFormValid = ({...rest}) => {
            let isValid=true

            Object.values(rest).forEach(element => {
                element === null && (isValid=false);
            });

        return isValid;
        }
        if (isFormValid(this.state)) {
            axios.get("http://www.geoplugin.net/json.gp")
            .then(response => this.setState({sourceipaddr:response}))
            const data = {
                Username: this.state.username,
                Password: this.state.password,
                SourceIpAddr: this.state.sourceIPAddr,
            };
            return fetch("http://127.0.0.1:8000/login", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            }).then(response => response.json())
            .then((jsonData) => {
                if (jsonData.statuscode === 200) {
                    this.setState({loggedin:true})
                } else if(jsonData.statuscode === 401) {
                    this.setState({unauth: true})
                }
            })
            .catch(error => {
                this.handleError(error);
            });
        } else {
            console.error("Empty Username/Password");
        }
}

    render() {
        if (this.state.loggedin) {
            return(
                <Messenger usernamePass={this.state.username}>Loading...</Messenger>
                <Redirect to="/messenger"/>
            );
        }
        return (
            <div className="base-container" ref={this.props.containerRef}>
                <div className="header">Gomessenger Login</div>
                <div className="content">
                    <div className="image">
                        <img src={LogoImage} alt=""/>
                    </div>
                    <div className="form">
                        <div className="form-group">
                            <label htmlFor="username">Username</label>
                            <input type="text" name="username" placeholder="username" onChange={this.handleChange}></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Password</label>
                            <input type="password" name="password" placeholder="password" onChange={this.handleChange}></input>
                        </div>
                    </div>
                </div>
                <div className="footer">
                    <button type="button" className ="btn" onClick={this.doLogin}> 
                        Login
                    </button> &nbsp;&nbsp;&nbsp;
                    <Link to="/signup">Sign-Up</Link>
                </div>
                {this.state.unauth && (
                    <div>
                        <h3>Incorrect Username/Password</h3>
                    </div>
                )} 
            </div> 
        );
    }
}

export default Login;