import axios from "axios";

class OrderClient {
    static async postOrder(item) {
        try {
            const create = await axios.post(
                "http://localhost:8080/api/order/",
                JSON.stringify(item),
                { headers: { "Content-Type": "application/json", "Authorization": localStorage.getItem("token") } }
            );
            return create.data;
        } catch (err) {
            if (err.response) {
                return err.response.data;
            }
            return { error: "Unexpected Error" };
        }
    }
}

export default OrderClient