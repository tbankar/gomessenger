import React from 'react';
import LogoImage from '../../login.svg';
import { Redirect } from 'react-router-dom';
import request from 'superagent';
import axios from 'axios'

// import { Base64 } from 'js-base64';


export class Register extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
                username:null,
                email:null,
                fullname:null,
                password:null,
                bLogin: false,
                response : {regResp: []},
                emptyFields: {
                    username:"",
                    email:"",
                    fullname:"",
                    password:"",
                },
                submitDisabled:true,
                validEmail:false,
                validUsername:false,
                validFullname:false,
                validPassword:false,
                createUserResp:"",
        };

        this.doSubmit = this.doSubmit.bind(this)
        this.backLogin = this.backLogin.bind(this)

    }

    handleChange = (event) => {
        const {name,value} = event.target

        let emptyFields = {...this.state.emptyFields};
        const validateEmail = RegExp(
            /^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/
          );

        switch(name) {
            case "fullname":
                    if (value.length < 5) {
                        emptyFields.fullname = "Fullname characters length should be > 5";
                        this.setState({validFullname:false})
                    } else {
                        emptyFields.fullname = ""
                        this.setState({validFullname:true})
                    }
                    break
            case "username":
                if (value.length < 5) {
                    emptyFields.username = "Username characters length should be > 5";
                        this.setState({validUsername:false})
                } else {
                    emptyFields.username = ""
                        this.setState({validUsername:true})
                    
                }
                break;
            case "password":
                    if (value.length < 1) {
                        this.setState({validPassword:false})
                        emptyFields.password = "Password should not be empty";
                    } else {
                        this.setState({validPassword:true})
                        emptyFields.password = ""
                    }
                break;
            case "email":
                if (!validateEmail.test(value)) {
                        this.setState({validEmail:false})
                        emptyFields.email = "Invalid Email";
                } else {
                    emptyFields.email = ""
                    this.setState({validEmail:true})
                }
                break;
            default:
                break
        }
        this.setState({submitDisabled: !(this.state.validEmail && this.state.validFullname && this.state.validPassword && this.state.validUsername)})
        this.setState({emptyFields, [name]: value});
    }

    handleResponseError(response) {
        throw new Error("HTTP error, status = " + response.status);
    }

    handleError(error) {
        console.log(error.message);
    }

    backLogin = (event) => {
        this.setState({bLogin:true})
    }

    doSubmit = (event) => {
        event.preventDefault()
        axios.get("http://www.geoplugin.net/json.gp")
        .then(response => this.setState({SourceIpAddr:response}))
        const data = {
            Username: this.state.username,
            Password: this.state.password,
            Fullname: this.state.fullname,
            Email: this.state.email,
            SourceIpAddr: this.state.SourceIpAddr
        };
        // data.Password = Base64.encode(data.Password)
        request
            .post('http://127.0.0.1:8000/create')
            .set('Content-Type', 'application/json')
            .send(JSON.stringify(data))
            .end(function(_err, res){
                    this.setState({createUserResp:res.text})
        });  
    }

    render() {
        const {emptyFields} = this.state;
        return ( 
            <div className="base-container" ref={this.props.containerRef}>
                <div className="header">Gomessenger Register</div>
                <div className="content">
                    <div className="image">
                        <img src={LogoImage} alt=""/>
                    </div>
                    <div className="form" onSubmit={this.doSubmit}>
                        <div className="form-group">
                            <label htmlFor="username">Username</label>
                            <input type="text" name="username" placeholder="username" onChange={this.handleChange} ></input>
                            {emptyFields.username.length >= 0 && (
                                <span className="errorMessage">{emptyFields.username}</span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Password</label>
                            <input type="password" name="password" placeholder="password" onChange={this.handleChange} ></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="email">Email</label>
                            <input type="text" name="email" placeholder="email" onChange = {this.handleChange}></input>
                            {emptyFields.email.length >= 0 && (
                                <span className="errorMessage">{emptyFields.email}</span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="name">Full Name</label>
                            <input type="text" name="fullname" placeholder="full name" onChange={this.handleChange}></input>
                            {emptyFields.fullname.length >= 0 && (
                                <span className="errorMessage">{emptyFields.fullname}</span>
                            )}
                        </div>
                    </div>
                </div>
                <div className="register">
                    <button type="submit" className="btn" onClick={this.backLogin}>Back to Login</button>&nbsp;
                    {this.state.bLogin && (
                        <Redirect to="/"></Redirect>
                    )}
                    <button disabled={this.state.submitDisabled} type="submit" className="btn" onClick={this.doSubmit}>Register Me</button>
                </div>
                <div className="createUserMsg">
                    {this.state.createUserResp !== "Success" && (
                        <h3>{this.state.createUserResp}</h3>
                    )}
                    {this.state.createUserResp === "Success" && (
                        <h3>Registration successful</h3>
                    )}
                </div>
            </div> 
        );
    }
}
