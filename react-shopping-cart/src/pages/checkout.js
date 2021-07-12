import React, {useContext, useEffect, useState} from 'react';
import { Link, useHistory } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux';
import CartClient from "../APIClients/CartClient";
import CheckoutClient from "../APIClients/CheckoutClient";
import OrderClient from "../APIClients/OrderClient";
import CheckoutSuccess from "./CheckoutSuccess";
import {CheckoutContext} from "../contexts/CheckoutContext";

function Checkout(props) {

  const [name, setName] = useState(null);
  const [email, setEmail] = useState(null);
  const [address, setAddress] = useState(null);
  const [ccnumber, setCreditNumber] = useState(null);
  const [expirationMonth, setExpirationMonth] = useState(null);
  const [expirationYear, setExpirationYear] = useState(null);
  const [cvv, setCVV] = useState(null);
  const history = useHistory();

  const submitHandler = async (e) => {
    e.preventDefault();
    
    OrderClient.postOrder(
        {
          email,
          name,
          "street_address": address,
          "items": JSON.parse(localStorage.getItem('cart'))
        });
    const result = await CheckoutClient.postCheckout(
        {
          name,
          "street_address": address,
          "card": {
            "number": ccnumber,
            "expiration_month": expirationMonth,
            "expiration_year": expirationYear,
            cvv
          }
        });

        if(result.succeeded){
          history.push('checkout-success');
        } else {
          history.push('checkout-failure')
        }
  }

    return <div>
      <div className="form">
        <form onSubmit={submitHandler}>
          <ul className="form-container">
            <li>
              <h2>Checkout</h2>
            </li>

            <li>
              <label htmlFor="name">
                Name
              </label>
              <input value={name} type="text" name="name" id="name" onChange={(e) => setName(e.target.value)}>
              </input>
            </li>

            <li>
              <label htmlFor="email">
                Email
              </label>
              <input value={email} type="text" name="email" id="email" onChange={(e) => setEmail(e.target.value)}>
              </input>
            </li>

            <li>
              <label htmlFor="address">
                Address
              </label>
              <input value={address} type="text" name="address" id="address" onChange={(e) => setAddress(e.target.value)}>
              </input>
            </li>

            <li>
              <label htmlFor="ccnumber">
                Credit Card Number
              </label>
              <input value={ccnumber} type="text" name="ccnumber" id="ccnumber"
                     onChange={(e) => setCreditNumber(e.target.value)}>
              </input>
            </li>

            <li>
              <label htmlFor="expirationMonth">
                Expiration Month
              </label>
              <input value={expirationMonth} type="text" name="expirationMonth" id="expirationMonth"
                     onChange={(e) => setExpirationMonth(e.target.value)}>
              </input>
            </li>

            <li>
              <label htmlFor="expirationYear">
                Expiration Year
              </label>
              <input value={expirationYear} type="text" name="expirationYear" id="expirationYear"
                     onChange={(e) => setExpirationYear(e.target.value)}>
              </input>
            </li>

            <li>
              <label htmlFor="cvv">
                CVV
              </label>
              <input value={cvv} type="text" name="cvv" id="cvv" onChange={(e) => setCVV(e.target.value)}>
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