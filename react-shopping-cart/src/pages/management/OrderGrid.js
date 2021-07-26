import React, { useContext, useState } from 'react';
import styles from '../../Styles/ProductsGrid.module.scss';
import OrderCard from './OrderCard'
import { OrdersContext } from '../../contexts/OrdersContext';

const OrderGrid = () => {
    console.log()
    const { orders } = useContext(OrdersContext);
    const [search, setSearch] = useState("");

    return (
        <div className={styles.p__container}>
            <div className="row">
                <div className="col-sm-8">
                    <div className="py-3">
                        {orders.length} Orders
                    </div>
                </div>
                <div className="col-sm-4">
                    <div className="form-group">
                        <input type="text" name="" onChange={(e) => setSearch(e.target.value)} placeholder="Search Orders by Username" className="form-control" id=""/>
                    </div>
                </div>
            </div>
            <div className={styles.p__grid}>

                {
                    orders.filter(order => order.user_name.toLowerCase().includes(search.toLowerCase())).map(order => (
                        <OrderCard key={order.id} order={order}/>
                    ))
                }

            </div>
            <div className={styles.p__footer}>

            </div>
        </div>
    );
}

export default OrderGrid;