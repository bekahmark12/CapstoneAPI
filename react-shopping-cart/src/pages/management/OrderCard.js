import React from 'react';
import { Card, Dropdown, Form, Button } from 'react-bootstrap';


const OrderCard = ({ order }) => {
    console.log(order)
    return (
        <Card style={{ width: '18rem' }}>
            <Card.Body>
                <Card.Title>Order from {order.user_email}</Card.Title>
                <Card.Subtitle className="mb-2 text-muted">Purchased by {order.user_name}</Card.Subtitle>
                    <br></br>
                    <label>Items Purchased:</label>
                    <ul>
                    {order.items_ordered.items_in_cart.map(oi=> (
                        <li>
                            <Card.Text>{oi.item.title} - {oi.item_quantity} </Card.Text>
                        </li>
                    ))}
                    </ul>
                    <br></br>
                <Dropdown
                    size="sm">
                    <Dropdown.Toggle variant="success" id="dropdown-basic" size="sm">
                        Shipping Method
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                        <Dropdown.Item href="#/action-1">USPS</Dropdown.Item>
                        <Dropdown.Item href="#/action-2">Fedex</Dropdown.Item>
                        <Dropdown.Item href="#/action-3"> UPS</Dropdown.Item>
                    </Dropdown.Menu>
                </Dropdown>
                <br></br>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Enter Tracking Number:</Form.Label>
                        <Form.Control type="email" placeholder="Tracking #" />
                    </Form.Group>

                    <Button variant="primary" type="submit" size="sm" align="center">
                        Submit
                    </Button>
                </Form>
            </Card.Body>
        </Card>
    )
}

export default OrderCard