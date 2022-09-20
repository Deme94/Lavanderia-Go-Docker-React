import React, { Component, Fragment } from 'react';
import { Link } from "react-router-dom";

export default class LavadorasView extends Component {

    constructor(props) {
        super(props);
        this.state = {
            lavadoras: [],
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
        fetch(`${process.env.REACT_APP_API_URL}/v1/washers`, requestOptions)
            .then((response) => response.json())
            .then((json) => {
                this.setState({
                    lavadoras: json.washers,
                    isLoaded: true,
                })
            })
    }

    render() {
        const { lavadoras, isLoaded } = this.state
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
                        {Object.keys(lavadoras).map((keyName, i) => (
                            lavadoras[keyName].status === "green" ? (
                                <tr key={lavadoras[keyName].id}>
                                    <th scope="row"><Link to={`/lavadora`} onClick={OpenLavadora(this.props, lavadoras[keyName])}>LAVADORA {lavadoras[keyName].id}</Link></th>
                                    <td>{lavadoras[keyName].status} <img width="20" height="20" src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Green_Light_Icon.svg/1200px-Green_Light_Icon.svg.png" alt="" /></td>
                                    <td>{lavadoras[keyName].price} €</td>
                                </tr>
                            ) : (
                                <tr key={lavadoras[keyName].id}>
                                    <th scope="row">LAVADORA {lavadoras[keyName].id}</th>
                                    <td>{lavadoras[keyName].status} <img width="20" height="20" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/1f/Red_Light_Icon.svg/2048px-Red_Light_Icon.svg.png" alt="" /></td>
                                    <td>{lavadoras[keyName].price} €</td>
                                </tr>
                            )
                        ))}
                    </tbody>
                </table>
            </Fragment>
        );
    }
}

function OpenLavadora(props, lavadora) {
    return (() => (props.handleMaquina(lavadora)))
}