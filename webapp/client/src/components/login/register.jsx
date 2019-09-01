import React from 'react';
import LogoImage from '../../login.svg';

export class Register extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        return ( 
            <div className="base-container" ref={this.props.containerRef}>
                <div className="header">Register</div>
                <div className="content">
                    <div className="image">
                        <img src={LogoImage} alt=""/>
                    </div>
                    <div className="form">
                        <div className="form-group">
                            <label htmlFor="username">Username</label>
                            <input type="text" name="username" placeholder="username"></input>
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
                            <input type="text" name="name" placeholder="name"></input>
                        </div>
                    </div>
                </div>
                <div className="footer">
                    <button type="button" className="btn">
                        Register
                    </button>
                </div>
            </div> 
        );
    }
}