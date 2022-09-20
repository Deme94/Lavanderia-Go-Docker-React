import React, { Component, Fragment } from 'react';
import { Link } from "react-router-dom";
import Input from "./input";
import Alert from "./alert";

export default class RegisterView extends Component {

    constructor(props) {
        super(props)
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)

        this.state = {
            email: "",
            password: "",
            confirmPassword: "",
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

        if (this.state.password.length < 8 || !/\d/.test(this.state.password)){
            errors.push("password")
        }

        if (this.state.password !== this.state.confirmPassword){
            errors.push("confirmPassword")
        }

        this.setState({
            errors: errors
        })

        if(errors.length > 0){
            return false
        }

        const data = new FormData(evt.target)
        const payload = Object.fromEntries(data.entries())

        const requestOptions = {
            method: "POST",
            body: JSON.stringify(payload),
        }

        fetch(`${process.env.REACT_APP_API_URL}/v1/register`, requestOptions)
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
                    this.props.history.push('/registerConfirmed')
                }
            })
    }

    handleJWTChange(jwt) {
        this.props.handleJWTChange(jwt)
    }

    hasError(key){
        return this.state.errors.indexOf(key) !== -1
    }

    render() {
        return (
            <Fragment>
                <h2>Registro</h2>
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
                        errorMsg={"Por favor introduzca un correo electrónico válido"}
                    />
                    <Input 
                        title={"Password"}
                        type={"password"}
                        name={"password"}
                        handleChange={this.handleChange}
                        className={this.hasError("password") ? "is-invalid" : ""}
                        errorDiv={this.hasError("password") ? "text-danger" : "d-none"}
                        errorMsg={"La contraseña debe tener al menos 8 caracteres y 1 dígito"}
                    />
                    <Input 
                        title={"Confirm password"}
                        type={"password"}
                        name={"confirmPassword"}
                        handleChange={this.handleChange}
                        className={this.hasError("confirmPassword") ? "is-invalid" : ""}
                        errorDiv={this.hasError("confirmPassword") ? "text-danger" : "d-none"}
                        errorMsg={"Las contraseñas no coinciden"}
                    />

                    <button className="btn btn-primary">Crear cuenta</button>
                </form>
                <hr />
                <Link to={`/login`}>Iniciar sesión</Link>
            </Fragment>
        );
    }
}
