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

    async getAllOrders(cb) {
        try {
            console.log('you hit the orders get request')
            const orders = await axios.get(
                "http://localhost:8080/api/order/", 
                { headers: {"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyVHlwZSI6MSwiRW1haWwiOiJ0ZXN0MkBnbWFpbC5jb20iLCJleHAiOjE2MjcwNjY2NzcsImlzcyI6InVzZXItc2VydmljZSJ9.8tm5ha5LFafqFbznnSYFiOGlN5Chdq5cNIlnPjpXk1Q" }});
            return cb(orders.data);
        } catch (err) {
            if (err.response) {
                return cb(err.response.data);
            }
            return cb({ error: "Unexpected Error" });
        }
    }
}

export default new OrderClient();