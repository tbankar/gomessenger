import React from 'react';
import LogoImage from '../../login.svg';

export class Login extends React.Component {

    constructor(props) {
        super(props);
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
                            <input type="text" name="username" placeholder="username"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Password</label>
                            <input type="password" name="password" placeholder="password"></input>
                        </div>
                    </div>
                </div>
                <div className="footer">
                    <button type="button" className ="btn"> 
                        Login
                    </button>
                </div>
            </div> 
        );
    }
}