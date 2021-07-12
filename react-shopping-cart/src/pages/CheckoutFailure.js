import React from 'react';
import Layout from '../components/Layout';

const CheckoutFailure = () => {

    return (
        <Layout title="Store" description="This is the Store page" >
            <div >
                <div className="text-center mt-5">
                    <h1>Shoot, Checkout Failed on the Backend!</h1>
                    <p>Check out the API error logs!</p>
                </div>
            </div>
        </Layout>
    );
}

export default CheckoutFailure;