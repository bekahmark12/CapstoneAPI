import axios from "axios";
 class ItemClient{

     async getAllItems() {
         try {
             const items = await axios.get("http://localhost:8080/api/items");
             if (items.data.error) {
                 return items.data;
             }
             return items.data;
         } catch (err) {
             if (err.response) {
                 return err.response.data;
             }
             return { error: "Unexpected Error" };
         }
     }

     async getItemByName(item_name) {
         try {
             const item = await axios.get(`http://localhost:8080/api/items/${item_name}`);
             return item.data;
         } catch (err) {
             if (err.response) {
                 return err.response.data;
             }
             return { error: "Unexpected Error" };
         }
     }

     async createItem(item) {
         try {
             const create = await axios.post(
                 "http://localhost:8080/api/items/create",
                 JSON.stringify(item),
                 { headers: { "Content-Type": "application/json" } }
             );
             return create.data;
         } catch (err) {
             if (err.response) {
                 return err.response.data;
             }
             return { error: "Unexpected Error" };
         }
     }

     async updateItemTitle(item_id, title) {
         try {
             const update = await axios.put(
                 `http://localhost:8080/api/items/update-title/${item_id}`,
                 JSON.stringify({ title }),
                 { headers: { "Content-Type": "application/json" } }
             );
             return update.data;
         } catch (err) {
             if (err.response) {
                 return err.response.data;
             }
             return { error: "Unexpected Error" };
         }
     }

     async updateItemDescription(item_id, description) {
         try {
             const update = await axios.put(
                 `http://localhost:8080/api/items/update-description/${item_id}`,
                 JSON.stringify({ description }),
                 { headers: { "Content-Type": "application/json" } }
             );
             return update.data;
         } catch (err) {
             if (err.response) {
                 return err.response.data;
             }
             return { error: "Unexpected Error" };
         }
     }

     async updateItemPrice(item_id, price) {
         try {
             const update = await axios.put(
                 `http://localhost:8080/items/update-price/${item_id}`,
                 JSON.stringify({ price }),
                 { headers: { "Content-Type": "application/json" } }
             );
             return update.data;
         } catch (err) {
             if (err.response) {
                 return err.response.data;
             }
             return { error: "Unexpected Error" };
         }
     }

 }
