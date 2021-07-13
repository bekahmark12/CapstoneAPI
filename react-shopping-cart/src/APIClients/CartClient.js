import axios from "axios";

class CartClient {

    // static async getKeys() {
    //     return localStorage.getItem("")
    // }

    static async postCart(item_id, quantity) {
        try {
            const create = await axios.post(
                `http://localhost:8080/api/cart/${item_id}?qty=${quantity}`,
                null,
                { headers: {"Authorization": localStorage.getItem("token") } }
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