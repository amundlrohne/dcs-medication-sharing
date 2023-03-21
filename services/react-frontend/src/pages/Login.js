import React from 'react'

import '../css/Login.css'

class Login extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            usernameValue: '',
            passwordValue: '',
            errorMessage: '',
            showError: false
        }
    }

    render() {
        const { errorMessage, showError } = this.state
        return (
            <div className='login-component'>
                <p id='title'>Healthcare Provider</p>
                <input className='login-input' id='username' placeholder='Username' onChange={e => this.updateUsernameValue(e)}></input>
                <input className='login-input' id='password' placeholder='Password' onChange={e => this.updatePasswordValue(e)} type='password'></input>

                <button className='login-button' onClick={this.loginClick}>Login</button>

                { showError && (
                    <p id='login-error'>{errorMessage}</p>
                ) }
            </div>
        )
    }

    async componentDidMount () {
        let cookieData = await getCookie()
        console.log(cookieData)

        if (cookieData.message === 'success') {
            window.location = '/inbox'
        }
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

    updateErrorMessage = (value, visibility) => {
        this.setState({
            errorMessage: value,
            showError: visibility
        })

        setTimeout(() => { this.updateErrorMessage("", false) }, 3000)
    }

    loginClick = async () => {

        if (this.state.usernameValue.length > 0 && this.state.passwordValue.length > 0){
            let url = "http://localhost:8280/health-provider/verify"

            try {
                let user_data = {
                    username: this.state.usernameValue,
                    password: this.state.passwordValue
                }

                let data = await postData(url, user_data)
                console.log(data.message)

                if (data.message === 'invalid') {
                    this.updateErrorMessage("Invalid login credentials!", true)
                }

                console.log(data)
                if (data.message === 'success') {
                    console.log("Logged in!")

                    let cookieData = await getCookie()
                    console.log(cookieData)
                }
            } catch {

            }
        } else {
            this.updateErrorMessage("Please fill in all fields!", true)
        }
    }
}

const postData = async (url, data) => {
    const response = await fetch(url, {
        method: "POST",
        mode: "cors",
        credentials: "include",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })

    return response.json()
}

const getCookie = async () => {
    let url = "http://localhost:8280/health-provider/current"
    const response = await fetch(url, {
        method: "GET",
        mode: "cors",
        credentials: "include"
    })

    return response.json()
}

const deleteCookie = async () => {
    let url = "http://localhost:8280/health-provider"
    const response = await fetch(url, {
        method: "DELETE",
        mode: "cors",
        credentials: "include"
    })

    return response.json()
}

export default Login
