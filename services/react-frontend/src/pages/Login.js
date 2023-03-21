import React from 'react'

import '../css/Login.css'

class Login extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            usernameValue: '',
            passwordValue: ''
        }
    }

    render() {
        return (
            <div className='login-component'>
                <p id='title'>Healthcare Provider</p>
                <input className='login-input' id='username' placeholder='Username' onChange={e => this.updateUsernameValue(e)}></input>
                <input className='login-input' id='password' placeholder='Password' onChange={e => this.updatePasswordValue(e)} type='password'></input>

                <button className='login-button' onClick={this.loginClick}>Login</button>
            </div>
        )
    }

    componentDidMount() {
        
    }

    updateUsernameValue = (e) => {
        const val = e.target.value
        this.setState({
            usernameValue: val
        })
    }

    updatePasswordValue = (e) => {
        const val = e.target.value
        this.setState({
            passwordValue: val
        })
    }

    loginClick = () => {
        console.log(this.state.usernameValue)
    }
}

export default Login
