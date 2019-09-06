import React from 'react';
import LogoImage from '../../login.svg';

export class Register extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
                username:null,
                email:null,
                fullname:null,
                password:null,
                response : {regResp: []},
        };

        this.doSubmit = this.doSubmit.bind(this)

    }

    handleChange = (event) => {
        const {name,value} = event.target
        this.setState({[name]: value});
    }

    doSubmit = (event) => {
        event.preventDefault()
        const data = {
            Username: this.state.username,
            Password: this.state.password,
            Fullname: this.state.fullname,
            Email: this.state.email,
        };
        return fetch("http://127.0.0.1:8000/create", {
            method: 'POST',
            mode: 'no-cors',
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data)
        }).then(response => {
            return response.json()
        }).then(json => {
            this.setState({response:json});
        });
    }

    render() {
        return ( 
            <div className="base-container" ref={this.props.containerRef}>
                <div className="header">Register</div>
                <div className="content">
                    <div className="image">
                        <img src={LogoImage} alt=""/>
                    </div>
                    <div className="form" onSubmit={this.doSubmit}>
                        <div className="form-group">
                            <label htmlFor="username">Username</label>
                            <input type="text" name="username" placeholder="username" onChange={this.handleChange} ></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Password</label>
                            <input type="password" name="password" placeholder="password" onChange={this.handleChange} ></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="email">Email</label>
                            <input type="text" name="email" placeholder="email" onChange = {this.handleChange}></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="name">Full Name</label>
                            <input type="text" name="fullname" placeholder="full name" onChange={this.handleChange}></input>
                        </div>
                    </div>
                </div>
                <div className="register">
                    <button type="submit" className="btn" onClick={this.doSubmit}>Register</button>
                </div>
            </div> 
        );
    }
}