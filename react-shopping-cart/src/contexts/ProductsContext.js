import React, { createContext, useState } from 'react';
import { dummyProducts } from '../services/dummy';
export const ProductsContext = createContext()

const ProductsContextProvider = ({children}) => {

    //API call to get all products here
    const [products] = useState(dummyProducts);

    return ( 
        <ProductsContext.Provider value={{products}} >
            //What is children? A Placeholder?
            { children }
        </ProductsContext.Provider>
     );
}
 
export default ProductsContextProvider;