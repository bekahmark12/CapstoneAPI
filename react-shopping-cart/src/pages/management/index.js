import React from 'react';
import Layout from '../../components/Layout';
import { Card } from 'react-bootstrap';

const Test = () => {
    
    return ( 
        <Layout title="test" description="your test page worked" >
            <div >
                <div className="text-center mt-5">
                    <h1>test</h1>
                    <p>This is the test Page.</p>
                </div>
            </div>
        </Layout>
     );
}
 
export default Test;