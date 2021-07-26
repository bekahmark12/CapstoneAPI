import React, { useContext, useState } from 'react';
import ProductItem from './ProductItem';
import { ProductsContext } from '../../contexts/ProductsContext';
import styles from '../../Styles/ProductsGrid.module.scss';

const ProductsGrid = () => {

    const { products} = useContext(ProductsContext)
    const [search, setSearch] = useState("");

    return (
        <div className={styles.p__container}>
            <div className="row">
                <div className="col-sm-8">
                    <div className="py-3">
                        {products.length} Products
                    </div>
                </div>
                <div className="col-sm-4">
                    <div className="form-group">
                        <input type="text" name="" onChange={(e) => setSearch(e.target.value)} placeholder="Search product" className="form-control" id=""/>
                    </div>
                </div>
            </div>
            <div className={styles.p__grid}>

                {
                    products.filter(product => product.title.toLowerCase().includes(search.toLowerCase())).map(product => (
                        <ProductItem key={product.id} product={product}/>
                    ))
                }

            </div>
            <div className={styles.p__footer}>

            </div>
        </div>
    );
}

export default ProductsGrid;