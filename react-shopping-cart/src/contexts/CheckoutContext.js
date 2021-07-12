import React, { createContext, useReducer } from 'react';
import {CheckoutReducer} from './CheckoutReducer';

export const CheckoutContext = createContext()

const storage = localStorage.getItem('cart') ? JSON.parse(localStorage.getItem('cart')) : [];
const initialState = { checkoutComplete: false };

const CheckoutContextProvider = ({children}) => {

    const [state, dispatch] = useReducer(CheckoutReducer, initialState)

    const handleCheckoutComplete = () => {
        console.log('CHECKOUT_COMPLETE', state);
        dispatch({type: 'CHECKOUT_COMPLETE'})
    }

    const contextValues = {
        handleCheckoutComplete,
        ...state
    }

    return (
        <CheckoutContext.Provider value={contextValues} >
            { children }
        </CheckoutContext.Provider>
    );
}

export default CheckoutContextProvider;
