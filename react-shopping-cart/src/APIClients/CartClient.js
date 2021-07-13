import axios from "axios";

class CartClient {
    static async postCart(item_id, quantity) {
        try {
            const create = await axios.post(
                `http://localhost:8080/api/cart/${item_id}?qty=${quantity}`,
                null,
                { headers: {"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjI2MTg3NTE1LCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.nu4kOswTftX9hL5BoE8GSJtmTteeLarRhvCbKfIrrss" } }
            );
            return {succeeded: true, data: create.data};
        } catch (err) {
            if (err.response) {
                console.log(err.response)
                return {succeeded: false, data: err.response};
            }
            console.log(err.response)
            return {succeeded: false, data: "unexpected error"};
        }
    }
}

export default CartClient;