import React from 'react';
import Layout from '../../components/Layout';
import OrderGrid from './OrderGrid';


const Test = () => {
    
    return ( 
        <Layout title="test" description="your test page worked" >
            <div >
                <div className="text-center mt-5">
                    <h1>Management</h1>
                    <p>This is the Management Page.</p>
                </div>
                <OrderGrid/>
            </div>
        </Layout>
     );
}
 
export default Test;