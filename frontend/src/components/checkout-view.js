import React, { Component, Fragment } from 'react';
import { useHistory } from 'react-router-dom';
import { PaymentElement } from '@stripe/react-stripe-js';
import { loadStripe } from '@stripe/stripe-js';
import { useStripe, useElements, Elements } from '@stripe/react-stripe-js';
import axios from 'axios';

// Make sure to call `loadStripe` outside of a component’s render to avoid
// recreating the `Stripe` object on every render.
const stripePromise = loadStripe(process.env.REACT_APP_STRIPE_PK);

export default class CheckoutView extends Component {

  constructor(props) {
    super(props);

    this.handlerMsg = this.handlerMsg.bind(this);
    this.handlerDisabled = this.handlerDisabled.bind(this);
    this.state = {
      idMaquina: props.maquina.id,
      amount: props.maquina.price,
      clientSecret: "",
      isLoaded: false,
      errorMsg: "",
      disabled: false
    }
  }
  componentDidMount() {
    const headers = {
      'Content-Type': 'application/json',
      'Authorization': "Bearer " + this.props.jwt
    }
    axios.post(
      `${process.env.REACT_APP_API_URL}/v1/create-payment-intent`, { amount: this.state.amount }, { headers: headers }) // PONER CANTIDAD REAL DE DINERO
      .then(res => {
        this.setState({
          clientSecret: res.data.paymentIntent.clientSecret,
          isLoaded: true,
        })
      })
  }

  handlerMsg(msg) {
    this.setState({
      errorMsg: msg,
    })
  }

  handlerDisabled(d) {
    this.setState({
      disabled: d,
    })
  }

  render() {
    if (this.state.isLoaded) {
      const options = {
        // passing the client secret obtained from the server
        clientSecret: this.state.clientSecret
      };

      const checkParams = {
        handlerMsg: this.handlerMsg,
        handlerDisabled: this.handlerDisabled,
        idMaquina: this.state.idMaquina,
        amount: this.state.amount,
        disabled: this.state.disabled,
        jwt: this.props.jwt
      }

      const formulario = (<Elements stripe={stripePromise} options={options}>
        <CheckoutForm {...checkParams} />
      </Elements>)

      if (this.state.errorMsg === "") {
        return (
          <Fragment>
            {formulario}
          </Fragment>
        );
      } else {
        return (
          <Fragment>
            {formulario}
            <p>{this.state.errorMsg}</p>
          </Fragment>
        );
      }
    }
    else {
      return null
    }
  }
}
function CheckoutForm(props) {
  const stripe = useStripe();
  const elements = useElements();
  const history = useHistory();

  const handleSubmit = async (event) => {
    // We don't want to let default form submission happen here,
    // which would refresh the page.
    event.preventDefault();

    props.handlerDisabled(true)
    props.handlerMsg("")

    if (!stripe || !elements) {
      return;
    }

    const result = await stripe.confirmPayment({
      elements,
      // confirmParams: {
      //   return_url: "http://localhost:3000/"+props.maquina.id+"/confirmed",
      // },
      redirect: "if_required"
    });

    if (result.error) {
      props.handlerMsg(result.error.message)
      props.handlerDisabled(false)
    } else {
      // Successful payment
      const headers = {
        'Content-Type': 'application/json',
        'Authorization': "Bearer " + props.jwt
      }
      const body = JSON.stringify(
        {
          id: props.idMaquina
        })
      axios.post(
        `${process.env.REACT_APP_API_URL}/v1/confirm-payment`, body, {headers: headers}) // ACTIVAR MAQUINA Y PONER POST
        .then(res => {
          console.log("Pago realizado! Maquina=", props.idMaquina)
          history.push(`/` + props.idMaquina + `/confirmed`)
        })
    }
  };
  return (
    <Fragment>
      {props.disabled ?
        <form onSubmit={handleSubmit}>
          <PaymentElement />
          <br></br>
          <p>Loading ...</p> {/* LOADING */}
        </form>
        : <form onSubmit={handleSubmit}>
          <PaymentElement />
          <br></br>
          <p>Cantidad a pagar: {props.amount} €</p>
          <button className="btn btn-primary">Pagar</button>
        </form>}
    </Fragment>
  )
};