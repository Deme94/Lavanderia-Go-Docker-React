import React, { Fragment } from 'react';
import { Link, useHistory } from "react-router-dom";

export default function LavadoraView(props) {
    var lavadora = props.maquina
    return (
        <Fragment>
            <h5>LAVADORA {lavadora.id}</h5>
            <p> {lavadora.price.toFixed(2).toString().replace(".", ",")} €</p>
            <Link to={`/checkout`}><button type="button" className="btn btn-primary">PAGAR</button></Link>
        </Fragment>
    );
}