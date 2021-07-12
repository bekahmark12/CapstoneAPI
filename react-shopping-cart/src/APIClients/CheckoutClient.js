import axios from "axios";

class CheckoutClient {
    // static async getKeys() {
    //     return localStorage.getItem("")
    // }

    static async postCheckout(item) {
        try {
            const create = await axios.post(
                "http://localhost:8080/api/checkout/",
                JSON.stringify(item),
                { headers: { "Content-Type": "application/json", "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjI2MTI5NzMyLCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.f9Orlj-ZszMIaX9mWFa5h6ZcsbSdbbqPi-YGNlHM7VY" } }
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