import React, { Fragment } from 'react';
import { Link } from "react-router-dom";

import BigButton from "./bigbutton.js";

export default function SelectionView(props) {
    return (
        <Fragment>
            <div className="row">
                <div className="col-sm">
                    <Link to={`/secadoras`}><BigButton title="SECAR" url="https://key0.cc/images/preview/2088625_0386581685a4b1aa8d737e519e52bb3b.png" /></Link>
                </div>
                <div className="col-sm">
                    <Link to={`/lavadoras`}><BigButton title="LAVAR" url="https://key0.cc/images/preview/2088625_0386581685a4b1aa8d737e519e52bb3b.png" /></Link>
                </div>
            </div>
        </Fragment>
    );
}

