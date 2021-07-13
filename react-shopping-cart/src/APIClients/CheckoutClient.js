import axios from "axios";

class CheckoutClient {
    static async postCheckout(item) {
        try {
            const create = await axios.post(
                "http://localhost:8080/api/checkout/",
                JSON.stringify(item),
                { headers: { "Content-Type": "application/json", "Authorization": localStorage.getItem("token") } }
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