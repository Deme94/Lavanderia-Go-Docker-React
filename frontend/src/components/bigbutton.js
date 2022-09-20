import React from 'react';

import "./bigbutton.css";

export default function BigButton(props) {
  return (
    <div className="mycard" >
      <img src={props.url} className="photo" alt="" />
      <div className="card-body">{props.title}</div>
    </div>
  );
}