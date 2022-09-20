import React, { Fragment } from 'react';
import { Link, useHistory } from "react-router-dom";

export default function SecadoraView(props) {
    var secadora = props.maquina
    return (
        <Fragment>
            <h5>SECADORA {secadora.id}</h5>
            <p> {secadora.price.toFixed(2).toString().replace(".", ",")} €</p>
            <Link to={`/checkout`}><button type="button" className="btn btn-primary">PAGAR</button></Link>
        </Fragment>
    );
}