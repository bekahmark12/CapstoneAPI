import React, {createContext, useEffect, useState} from 'react';
import ItemClient from '../APIClients/ItemClient'

export const ProductsContext = createContext() ;


const ProductsContextProvider = ({children}) => {

    const [products, setProducts] = useState(null);

    useEffect(() => {
        let mounted = true;
        if (mounted) {
            ItemClient.getAllItems((data) => {
                setProducts(data);
            })
        }
        return () => mounted = false;
    }, [])

    if(!products){
        return <h1>Loading...</h1>
    }



    return ( 
        <ProductsContext.Provider value={{products}} >
            { children }
        </ProductsContext.Provider>
     );
}
 
export default ProductsContextProvider;