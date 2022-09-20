import React, { Component, Fragment } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';

import FacebookLogin from 'react-facebook-login';
import GoogleLogin from 'react-google-login';

import Input from "./input";
import Alert from "./alert";

export default class LoginView extends Component {

    constructor(props) {
        super(props)
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)

        this.state = {
            email: "",
            password: "",
            error: null,
            errors: [],
            alert: {
                type: "d-none",
                message: "",
            }
        }
    }

    componentDidMount() {

    }

    handleChange(evt) {
        let value = evt.target.value
        let name = evt.target.name
        this.setState((prevState) => ({
            ...prevState,
            [name]: value,
        }))
    }

    handleSubmit = (evt) => {
        evt.preventDefault()

        let errors = []
        if (this.state.email === "") {
            errors.push("email")
        }

        if (this.state.password === "") {
            errors.push("password")
        }

        this.setState({
            errors: errors
        })

        if (errors.length > 0) {
            return false
        }

        const data = new FormData(evt.target)
        const payload = Object.fromEntries(data.entries())

        const requestOptions = {
            method: "POST",
            body: JSON.stringify(payload),
        }

        fetch(`${process.env.REACT_APP_API_URL}/v1/login`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    this.setState({
                        alert: {
                            type: "alert-danger",
                            message: data.error.message,
                        }
                    })
                } else {
                    console.log(data)
                    this.handleJWTChange(Object.values(data)[0])
                }
            })
    }

    handleJWTChange(jwt) {
        this.props.handleJWTChange(jwt)
    }

    hasError(key) {
        return this.state.errors.indexOf(key) !== -1
    }

    render() {
        const responseFacebook = (response) => {
            console.log(response);
        }

        const responseGoogle = (response) => {
            console.log(response.Ru.Hv);
            const email = response.Ru.Hv
            const token = response.tokenId
            const googleId = response.googleId
            const headers = {
                'Content-Type': 'application/json',
            }
            const body = JSON.stringify(
                {
                    email: email,
                    token: token,
                    googleId: googleId
                })
            axios.post(
                `${process.env.REACT_APP_API_URL}/v1/login-google`, body, { headers: headers }) // ACTIVAR MAQUINA Y PONER POST
                .then((data) => {
                    if (data.error) {
                        this.setState({
                            alert: {
                                type: "alert-danger",
                                message: data.error.message,
                            }
                        })
                    } else {
                        console.log('Login Google validado')
                        console.log(Object.values(data)[0])
                        this.handleJWTChange(Object.values(data)[0].response)
                    }
                })
        }
        return (
            <Fragment>
                <h2>Inicio de sesión</h2>
                <hr />
                <Alert
                    alertType={this.state.alert.type}
                    alertMessage={this.state.alert.message}
                />

                <form className="pt-3" onSubmit={this.handleSubmit}>
                    <Input
                        title={"Email"}
                        type={"email"}
                        name={"email"}
                        handleChange={this.handleChange}
                        className={this.hasError("email") ? "is-invalid" : ""}
                        errorDiv={this.hasError("email") ? "text-danger" : "d-none"}
                        errorMsg={"Please enter a valid email address"}
                    />
                    <Input
                        title={"Password"}
                        type={"password"}
                        name={"password"}
                        handleChange={this.handleChange}
                        className={this.hasError("password") ? "is-invalid" : ""}
                        errorDiv={this.hasError("password") ? "text-danger" : "d-none"}
                        errorMsg={"Please enter a password"}
                    />

                    <button className="btn btn-primary">Iniciar sesión</button>
                </form>
                <hr />
                <Link to={`/register`}>Registrar cuenta nueva</Link>


                {/* <FacebookLogin
                    appId="" //APP ID NOT CREATED YET
                    fields="name,email,picture"
                    callback={responseFacebook}
                /> */}
                <br />
                <br />

                <GoogleLogin
                    clientId="866281508187-3t47rjldun9uqdbk633pb61pgnf755gd.apps.googleusercontent.com" //CLIENTID NOT CREATED YET
                    buttonText="LOGIN WITH GOOGLE"
                    onSuccess={responseGoogle}
                    onFailure={responseGoogle}
                />
            </Fragment >
        );
    }
}
