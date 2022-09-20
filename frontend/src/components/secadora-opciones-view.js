import React, { Fragment } from "react";
import { Link, useHistory } from "react-router-dom";

export default function SecadoraOpcionesView(props) {
  var secadora = props.maquina
  return (
    <Fragment>
      <h5>SECADORA {secadora.id}</h5>
      <table className="table table-bordered">
        <tbody>
          <tr key="1">
            <td>5 min</td>
            <td>1 €</td>
            <td scope="row"><Link to={`/secadora`} onClick={() => SelectOpcionSecadora(secadora, 1)}><button type="button" className="btn btn-primary">PAGAR</button></Link></td>
          </tr>
          <tr key="2">
            <td>10 min</td>
            <td>2 €</td>
            <td scope="row"><Link to={`/secadora`} onClick={() => SelectOpcionSecadora(secadora, 2)}><button type="button" className="btn btn-primary">PAGAR</button></Link></td>
          </tr>
          <tr key="3">
            <td>15 min</td>
            <td>3 €</td>
            <td scope="row"><Link to={`/secadora`} onClick={() => SelectOpcionSecadora(secadora, 3)}><button type="button" className="btn btn-primary">PAGAR</button></Link></td>
          </tr>
          <tr key="4">
            <td>20 min</td>
            <td>4 €</td>
            <td scope="row"><Link to={`/secadora`} onClick={() => SelectOpcionSecadora(secadora, 4)}><button type="button" className="btn btn-primary">PAGAR</button></Link></td>
          </tr>
        </tbody>
      </table>
    </Fragment>
  );
};

function SelectOpcionSecadora(secadora, precio) {
  secadora.price = precio
}