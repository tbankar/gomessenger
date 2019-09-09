import React from 'react';
import LogoImage from '../../login.svg';

export class Login extends React.Component {

    constructor(props) {
        super(props);

        this.state ={
            username:null,
            password:null,
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
        console.log(this.password)
        const data = {
            Username: this.state.username,
            Password: this.state.password,
        };
        return fetch("http://127.0.0.1:8000/login", {
            method: 'POST',
            mode: 'no-cors',
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        }).then(response => {
            if (!response.ok) {
                this.handleResponseError(response);
            }
        }).catch(error => {
            this.handleError(error);
        });
    }

    render() {
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
                    </button>
                </div>
            </div> 
        );
    }
}