import React, {Component} from 'react';
import ReactDOM from 'react-dom';

import "./index.css";

import Content from "./content.js";

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

class App extends Component{
    render(){
      return(
        <div>
          {/* <div className="header">
            <div className="row">
                <div className="col-sm-12">
                  <img className="logo" src="https://1000marcas.net/wp-content/uploads/2020/02/Google-logo.jpg" alt=""/>
                </div>
            </div>
          </div> */}
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <Content />
        </div>
      );
    }
}

ReactDOM.render(<App />, document.getElementById('root'));