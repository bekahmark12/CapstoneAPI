import axios from "axios";
class CartClient{

    static postItemToCart = async(itemId, itemQty) =>{
        try{
            const request = await axios.post(
                `http://localhost:8080/api/cart/${itemId}qty=${itemQty}`,
                { headers: { "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjI2MTI5NzMyLCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.f9Orlj-ZszMIaX9mWFa5h6ZcsbSdbbqPi-YGNlHM7VY" } }

            )
            return {succeeded: true, data: request.data};
        } catch (err) {
            if (err.response) {
                return {succeeded: false, data: err.response};
            }
            return {succeeded: false, data: "unexpected error :("};
        }
    }


}
export default CartClient
