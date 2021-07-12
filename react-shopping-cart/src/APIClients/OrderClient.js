import axios from "axios";

class OrderClient {
    static async postOrder(item) {
        try {
            const create = await axios.post(
                "http://localhost:8080/api/order/",
                JSON.stringify(item),
                { headers: { "Content-Type": "application/json", "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjI2MTI5NzMyLCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.f9Orlj-ZszMIaX9mWFa5h6ZcsbSdbbqPi-YGNlHM7VY" } }
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