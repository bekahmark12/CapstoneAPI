import React, {createContext, useEffect, useState} from 'react';
import { dummyProducts } from '../services/dummy';
import ItemClient from '../APIClients/ItemClient'
import axios from "axios";
export const ProductsContext = createContext() ;


const ProductsContextProvider = ({children}) => {

    // async function fetchItems() {
    //     return ItemClient.getAllItems();
    // }
    //const [products] = useState(ItemClient.getAllItems);
    const [products] = useState(dummyProducts)

    return ( 
        <ProductsContext.Provider value={{products}} >
            //What is children? A Placeholder?
            { children }
        </ProductsContext.Provider>
     );
}
 
export default ProductsContextProvider;