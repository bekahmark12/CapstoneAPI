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

const Routes = () => {
  return (
    <Router>
        <Switch>
          <Route path="/about" component={About} />
          <Route exact path="/" component={Store}/>
          <Route path="/cart" component={Cart} />
          <Route path="/checkout" component={Checkout} />
          <Route path="*" component={NotFound} />
        </Switch>
    </Router>
  );
}

export default Routes;