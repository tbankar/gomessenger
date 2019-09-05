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
        };

        this.handleFormSubmit = this.handleFormSubmit.bind(this)

    }
    handleChange = e => {
        const {name,value} = e.target
        this.setState({ [name]: value,[name]:value });
    }

    handleFormSubmit = e => {
        e.preventDefault()
        console.log(this.state.username)
        console.log(this.state.password)
    }


    render() {
        return ( 
            <div className="base-container" ref={this.props.containerRef}>
                <div className="header">Register</div>
                <div className="content">
                    <div className="image">
                        <img src={LogoImage} alt=""/>
                    </div>
                    <div className="form" onSubmit={this.handleFormSubmit}>
                        <div className="form-group">
                            <label htmlFor="username">Username</label>
                            <input type="text" name="username" placeholder="username" onChange={this.handleChange} ></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Password</label>
                            <input type="password" name="password" placeholder="password"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="email">Email</label>
                            <input type="text" name="email" placeholder="email"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="name">Full Name</label>
                            <input type="text" name="name" placeholder="full name"></input>
                        </div>
                    </div>
                </div>
                <div className="register">
                    <button type="submit" className="btn" onClick={this.handleFormSubmit}>Register</button>
                </div>
            </div> 
        );
    }
}