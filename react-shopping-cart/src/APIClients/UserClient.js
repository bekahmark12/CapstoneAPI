import axios from "axios";

class UserClient {
    async getBearerToken(cb) {
        try {
            const credentials = await axios.get("http://localhost:8080/users/login/");
            return cb(credentials.data);
        } catch (err) {
            if (err.response) {
                return cb(err.response.data);
            }
            return cb({ error: "Unexpected Error" });
        }
    }
}