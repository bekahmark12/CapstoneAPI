
const Storage = (cartItems) => {
    localStorage.setItem('cart', JSON.stringify(cartItems.length > 0 ? cartItems: []));
}

export const CheckoutReducer = (state, action) => {
    switch (action.type) {
        case "CHECKOUT_COMPLETE":
            return {
                checkout: true,
            }
        default:
            return state
    }
}