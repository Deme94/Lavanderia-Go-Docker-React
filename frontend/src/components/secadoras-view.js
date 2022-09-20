import React, { Component, Fragment } from 'react';
import { Link } from "react-router-dom";

export default class SecadorasView extends Component {

    constructor(props) {
        super(props);
        this.state = {
            secadoras: [],
            isLoaded: false,
        }
    }
    componentDidMount() {
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': "Bearer " + this.props.jwt
        }
        const requestOptions = {
            method: "GET",
            headers: headers,
        }
        fetch(`${process.env.REACT_APP_API_URL}/v1/dryers`, requestOptions)
            .then((response) => response.json())
            .then((json) => {
                this.setState({
                    secadoras: json.dryers,
                    isLoaded: true,
                })
            })
    }

    render() {
        const { secadoras, isLoaded } = this.state
        if (!isLoaded) {
            return (<p>Loading...</p>) // LOADING
        }
        return (
            <Fragment>
                <table className="table table-bordered">
                    {/* <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Estado</th>
                            <th scope="col">Precio</th>
                        </tr>
                    </thead> */}
                    <tbody>
                        {Object.keys(secadoras).map((keyName, i) => (
                            secadoras[keyName].status === "green" ? (
                                <tr key={secadoras[keyName].id}>
                                    <th scope="row"><Link to={`/secadora_opciones`} onClick={OpenSecadora(this.props, secadoras[keyName])}>SECADORA {secadoras[keyName].id}</Link></th>
                                    <td>{secadoras[keyName].status} <img width="20" height="20" src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Green_Light_Icon.svg/1200px-Green_Light_Icon.svg.png" alt="" /></td>
                                    <td>{secadoras[keyName].price} €</td>
                                </tr>
                            ) : (
                                <tr key={secadoras[keyName].id}>
                                    <th scope="row">SECADORA {secadoras[keyName].id}</th>
                                    <td>{secadoras[keyName].status} <img width="20" height="20" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/1f/Red_Light_Icon.svg/2048px-Red_Light_Icon.svg.png" alt="" /></td>
                                    <td>{secadoras[keyName].price} €</td>
                                </tr>
                            )
                        ))}
                    </tbody>
                </table>
            </Fragment>
        );
    }
}

function OpenSecadora(props, secadora) {
    return (() => (props.handleMaquina(secadora)))
}