import React, { Fragment } from 'react';
import { useParams } from "react-router-dom";

export default function ConfirmedView(props) {
    let { id } = useParams();
    return (
        <Fragment>
            <h2>Pago realizado con éxito.</h2>
            <br></br>
            <br></br>
            <h1>Pulse el botón START de la Máquina {id}</h1>
        </Fragment>
    );
}

