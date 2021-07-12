import React from 'react';
import Layout from '../components/Layout';

const CheckoutSuccess = () => {

    return (
        <Layout title="Store" description="This is the Store page" >
            <div >
                <div className="text-center mt-5">
                    <h1>Success!</h1>
                    <p>You have been successfully checked out!</p>
                </div>
            </div>
        </Layout>
    );
}

export default CheckoutSuccess;