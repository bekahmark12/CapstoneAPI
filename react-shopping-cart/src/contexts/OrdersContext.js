import React, {createContext, useEffect, useState} from 'react';
import OrderClient from '../APIClients/OrderClient'
export const OrdersContext = createContext() ;


const OrdersContextProvider = ({children}) => {

    const [orders, setOrders] = useState(null);

    useEffect(() => {
        let mounted = true;
        if (mounted) {
            OrderClient.getAllOrders((data) => {
                setOrders(data);
            })
        }
        return () => mounted = false;
    }, [])

    if(!orders){
        return <h1>Loading...</h1>
    }



    return ( 
        <OrdersContext.Provider value={{orders}} >
            { children }
        </OrdersContext.Provider>
     );
}
 
export default OrdersContextProvider;