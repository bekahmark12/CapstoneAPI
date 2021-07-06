import React, { useContext } from 'react';
import Layout from '../../components/Layout';
import { Form, Button } from 'react-bootstrap';
import { formatNumber } from '../../helpers/utils';
import { Link } from 'react-router-dom';

const AddProduct = () => {

    return (
        <Layout title="Cart" description="This is the Cart page" >
            <div >
                <div className="text-center mt-5">
                    <h1>New Product</h1>
                    <p>Add A New Product to the Current Selection.</p>
                    <Form>
                        <Form.Group controlId="formProductName">
                            <Form.Label>Product Name</Form.Label>
                            <Form.Control type="text" placeholder="Product Name" />
                        </Form.Group>

                        <Form.Group controlId="formProductDescription">
                            <Form.Label>Product Description</Form.Label>
                            <Form.Control type="text" placeholder="Description" />
                        </Form.Group>
                        <Form.Group controlId="formProductImageURL">
                            <Form.Label>Product Image</Form.Label>
                            <Form.Control type="text" placeholder="Image URL (file path for beta purposes)" />
                        </Form.Group>
                        <Form.Group controlId="formProductPrice">
                            <Form.Label>Price</Form.Label>
                            <Form.Control type="currency" placeholder="00.00" />
                        </Form.Group>
                        <Button variant="primary" type="submit">
                            Create Product
                        </Button>
                    </Form>
                </div>
            </div>
        </Layout>
    );
}

export default AddProduct;