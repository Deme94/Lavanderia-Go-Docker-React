import React, { Component } from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect
} from "react-router-dom";

import SelectionView from "./components/selection-view.js";
import LavadorasView from "./components/lavadoras-view.js";
import SecadorasView from "./components/secadoras-view.js";
import SecadoraOpcionesView from "./components/secadora-opciones-view.js";
import LavadoraView from "./components/lavadora-view.js";
import SecadoraView from "./components/secadora-view.js";
import CheckoutView from "./components/checkout-view.js";
import ConfirmedView from "./components/confirmed-view.js";
import LoginView from "./components/login-view.js";
import RegisterView from "./components/register-view.js";
import RegisterConfirmedView from "./components/register-confirmed-view.js";

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

export default class Content extends Component {

    constructor(props) {
        super(props)
        this.handleJWTChange = this.handleJWTChange.bind(this)
        this.handleMaquina = this.handleMaquina.bind(this)

        this.state = {
            jwt: "",
            maquina: null,
        }
    }

    handleJWTChange(jwt) {
        this.setState({
            jwt: jwt
        })
    }

    handleMaquina(m) {
        this.setState({
            maquina: m,
        })
    }

    render() {
        return (
            <Router>
                <Switch>
                    <Route path="/registerConfirmed" component={(props) =>
                        <div className="content">
                            <RegisterConfirmedView />
                        </div>
                    } />
                    <Route exact path="/register" component={(props) =>
                        <div className="content">
                            <RegisterView {...props} />
                        </div>
                    } />
                    <Route exact path="/login" component={(props) =>
                        this.state.jwt !== "" ? <Redirect to="/" /> :
                            <div className="content">
                                <LoginView handleJWTChange={this.handleJWTChange} />
                            </div>
                    } />
                    <Route path="/:id/confirmed" component={(props) =>
                        <div className="content">
                            <ConfirmedView />
                        </div>
                    } />
                    <Route path="/checkout" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <CheckoutView maquina={this.state.maquina} jwt={this.state.jwt} />
                            </div>
                    } />

                    <Route path="/secadora" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <SecadoraView maquina={this.state.maquina} />
                            </div>
                    } />
                    <Route path="/lavadora" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <LavadoraView maquina={this.state.maquina} />
                            </div>
                    } />
                    <Route path="/secadora_opciones" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <SecadoraOpcionesView maquina={this.state.maquina} />
                            </div>
                    } />
                    <Route path="/secadoras" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <SecadorasView handleMaquina={this.handleMaquina} jwt={this.state.jwt} />
                            </div>
                    } />
                    <Route path="/lavadoras" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <LavadorasView handleMaquina={this.handleMaquina} jwt={this.state.jwt} />
                            </div>
                    } />
                    <Route path="/" component={(props) =>
                        this.state.jwt === "" ? <Redirect to="/login" /> :
                            <div className="content">
                                <SelectionView />
                            </div>
                    } />
                </Switch>
            </Router>
        );
    }
}