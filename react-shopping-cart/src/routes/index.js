import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

import Store from '../pages/store';
import About from '../pages/About';
import NotFound from '../pages/NotFound';
import Cart from "../pages/cart";
import Checkout from "../pages/checkout";
import AddProduct from "../pages/addProduct";
import CheckoutSuccess from "../pages/CheckoutSuccess";
import CheckoutFailure from "../pages/CheckoutFailure"
import Login from "../pages/Login"
import Management from "../pages/Management";

const Routes = () => {
  return (
    <Router>
        <Switch>
          <Route path="/add-product" component={AddProduct} />
          <Route path="/about" component={About} />
          <Route exact path="/" component={Store}/>
          <Route path="/cart" component={Cart} />
          <Route path="/checkout" component={Checkout} />
          <Route path="/checkout-success" component={CheckoutSuccess}/>
          <Route path="/checkout-failure" component={CheckoutFailure}/>
          <Route path="/login" component={Login}/>
          <Route path="*" component={NotFound} />
          <Route path="/management" component={Management}/>
        </Switch>
    </Router>
  );
}

export default Routes;