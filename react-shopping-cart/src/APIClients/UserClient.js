import axios from "axios";

class UserClient {
    static async getBearerToken(credentials) {
        
        try {
            const token = await axios.post(
                "http://localhost:8080/api/users/",
                JSON.stringify(credentials),
                { headers: { "Content-Type": "application/json"}});
            return token.data;
        } catch (err) {
            if (err.response) {
                console.log(err.response)
                return err.response.data;
            }
            console.log(err.response)
            return { error: "Unexpected Error" };
        }
    }

    static async getUser() {
        try {
            const user = await axios.get(
                "http://localhost:8080/api/users/",
                { headers: { "Content-Type": "application/json", "Authorization": localStorage.getItem("token") }});
                return user.data;
        } catch (err) {
            if (err.response) {
                console.log(err.response)
                return err.response.data;
            }
            console.log(err.response)
            return { error: "Unexpected Error" };
        }
    }
}
export default UserClient