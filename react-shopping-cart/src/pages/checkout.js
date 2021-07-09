import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux';
//import CheckoutSteps from '../components/CheckoutSteps';

function Checkout(props) {

  const [name, setName] = useState(null);
  const [email, setEmail] = useState(null);
  const [address, setAddress] = useState(null);
  const [ccnumber, setCreditNumber] = useState(null);
  const [expiration, setExpiration] = useState(null);
  const [cvv, setCVV] = useState(null);

  const dispatch = useDispatch();
  const submitHandler = (e) => {
    e.preventDefault();
    props.history.push('payment');
  }
  return <div>
    <div className="form">
      <form onSubmit={submitHandler} >
        <ul className="form-container">
          <li>
            <h2>Checkout</h2>
          </li>

          <li>
            <label htmlFor="name">
              Name
          </label>
            <input type="text" name="name" id="name" onChange={(e) => setName(e.target.value)}>
            </input>
          </li>
          <li>
            <label htmlFor="email">
              Email
          </label>
            <input type="text" name="email" id="email" onChange={(e) => setEmail(e.target.value)}>
            </input>
          </li>      
          <li>
            <label htmlFor="address">
              Address
          </label>
            <input type="text" name="address" id="address" onChange={(e) => setAddress(e.target.value)}>
            </input>
          </li>
          <li>
            <label htmlFor="ccnumber">
              Credit Card Number
          </label>
            <input type="text" name="ccnumber" id="ccnumber" onChange={(e) => setCreditNumber(e.target.value)}>
            </input>
          </li>
          <li>
            <label htmlFor="expiration">
              Expiration Date
          </label>
            <input type="text" name="expiration" id="expiration" onChange={(e) => setExpiration(e.target.value)}>
            </input>
          </li>
          <li>
            <label htmlFor="cvv">
              CVV
          </label>
            <input type="text" name="cvv" id="cvv" onChange={(e) => setCVV(e.target.value)}>
            </input>
          </li>


          <li>
            <button type="submit" className="button primary">Continue</button>
          </li>

        </ul>
      </form>
    </div>
  </div>

}
export default Checkout;