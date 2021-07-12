import axios from "axios";

class CheckoutClient {
    static async postCheckout(item) {
        try {
            const create = await axios.post(
                "http://localhost:8080/api/checkout/",
                JSON.stringify(item),
                { headers: { "Content-Type": "application/json", "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjI2MTg3NTE1LCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.nu4kOswTftX9hL5BoE8GSJtmTteeLarRhvCbKfIrrss" } }
            );
            return {succeeded: true, data: create.data};
        } catch (err) {
            if (err.response) {
                return {succeeded: false, data: err.response};
            }
            return {succeeded: false, data: "unexpected error :("};
        }
    }
}

export default CheckoutClient