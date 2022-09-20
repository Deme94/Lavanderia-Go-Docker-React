import React, { Fragment } from 'react';
import { useParams } from "react-router-dom";

export default function RegisterConfirmedView() {
    let { id } = useParams();
    return (
        <Fragment>
            <h2>Cuenta registrada con éxito.</h2>
            <br></br>
            <br></br>
            <h1>Vuelva a recargar la página e inicie sesión. {id}</h1>
        </Fragment>
    );
}

