import React from 'react';
import Layout from '../components/Layout';

const Management = () => {

    return (
        <Layout title="Management" description="This is the Management page" >
            <div >
                <div className="text-center mt-5">
                    <h1>Orders</h1>
                    <p>(List of orders here)</p>
                </div>
                <div>
                    <form>
                        <input>USPS, UPS, FedEx</input>
                    </form>
                </div>
            </div>
        </Layout>
    );
}

export default Management;