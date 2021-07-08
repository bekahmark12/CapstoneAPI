# API ROUTES



## ITEM SERVICE
BASE URL: http://localhost:8080/api/items/

POST / : Inserts Item into database. Returns 204 on success. Returns 400 if unable to insert.

PUT /id:[0-9] : Updates Item in database based on the Item sent in request body. Returns 202 and updated item on success.<br/> Returns 400 if unable to update in database. Returns 404 if unable to find Item

GET / : Returns all Items stored. 

GET /id:[0-9] : Returns Item with the provided id. Returns 404 if Item is not found

DELETE /id:[0-9] : Deletes the Item with the provided id. Returns 404 if the item can not be found


POST and PUT endpoints have middleware that validate that a well structured Item has been provided in the request body. 

The middleware function will return a 400 if it is unable to read the json or if the Item fails the validation test.

Look below for an example request. Only the title and price fields are required.
```json
{
    "title":"poop",
    "price":2.4
}
```

## CART SERVICE

BASE URL: http://localhost:8080/api/cart/

POST /id:[0-9]?qty=[0-9] : Gets Item with the provided id and stores it with the provided quantity in a redis cache. Returns 204 on success.<br/> Returns 500 if unable to establish connection to Item Service and 400 if unable to store provided item in cache

GET / : Gets the current cart for the user. Returns 500 if unable to retrieve cart

DELETE / : Clears current users cart. Returns 202 on success. Returns 500 if unable to clear

DELETE /id:[0-9] : removes the item with the provided ID from the cart. Returns 400 if unable to remove item.

PATCH /id:[0-9]?qty=[0-9]: Updates items quantity in the current users cart.<br/>
returns 202 on success and 404 on failure

All endpoints in this service require a valid JWT provided by the User service<br/>
If no token is provided a 403 is returned. If the provided token is invalid a 401<br/>
is returned. 500 may be returned if Cart service is unable to reach User service

## ORDER SERVICE

BASE URL: http://localhost:8080/api/order/

POST / : Submits current users order and stores it in mongo instance. Returns<br/>
202 on success. Returns 400 if unable to complete order. May return 500 if unable<br/>
to reach User Service

All endpoints in this service require a valid JWT provided by the User service<br/>
If no token is provided a 403 is returned. If the provided token is invalid a 401<br/>
is returned. 500 may be returned if Cart service is unable to reach User service

## USER SERVICE

BASE URL: http://localhost:8080/api/users/

POST /login : validates user and returns token to use in other request.

POST /sign-up : creates user and stores 
