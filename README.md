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

## INSERT SCRIPT

INSERT SCRIPT : 

insert into items (id, image_url, title, description, price) values (1, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/portrait-2194457_1920_rnytv9.jpg', 'Matsoft', 'Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat.', 547.6);

insert into items (id, image_url, title, description, price) values (2, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/bicycle-1834265_1920_litf3x.jpg', 'Konklux', 'Suspendisse potenti. In eleifend quam a odio. In hac habitasse platea dictumst.
Maecenas ut massa quis augue luctus tincidunt. Nulla mollis molestie lorem. Quisque ut erat.', 245.0);

insert into items (id, image_url, title, description, price) values (3, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/quickdraws-6154461_1920_qnv6ga.jpg', 'Alpha', 'Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.
Etiam vel augue. Vestibulum rutrum rutrum neque. Aenean auctor gravida sem.', 668.4);

insert into items (id, image_url, title, description, price) values (4, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/person-690547_1920_u4n1fj.jpg', 'Tresom', 'Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue. Aliquam erat volutpat.In congue. Etiam justo. Etiam pretium iaculis justo.
In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus.', 981.1);

insert into items (id, image_url, title, description, price) values (5, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/bottles-774466_1920_d5u48q.jpg', 'Redhold', 'Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.
Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est. Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros.
Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue. Aliquam erat volutpat.', 25.7);

insert into items (id, image_url, title, description, price) values (6, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/stew-750846_1920_qdgtea.jpg', 'Bamity', 'In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus.
Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi.
Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque.', 408.0);

insert into items (id, image_url, title, description, price) values (7, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/fly-fishing-1149502_1920_j9819q.jpg', 'Fixflex', 'Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus.
Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.', 230.2);

insert into items (id, image_url, title, description, price) values (8, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/hiking-1149891_1920_j0gxw5.jpg', 'Regrant', 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Proin risus. Praesent lectus.
Vestibulum quam sapien, varius ut, blandit non, interdum in, ante. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Duis faucibus accumsan odio. Curabitur convallis.', 948.2);

insert into items (id, image_url, title, description, price) values (9, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/tent-548022_1920_kn2iiq.jpg', 'Lotlux', 'Maecenas ut massa quis augue luctus tincidunt. Nulla mollis molestie lorem. Quisque ut erat.', 670.6);

insert into items (id, image_url, title, description, price) values (10, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/rosemary-potatoes-1446677_1920_u6rudl.jpg', 'Zoolab', 'Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo. In blandit ultrices enim. Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
Proin interdum mauris non ligula pellentesque ultrices. Phasellus id sapien in sapien iaculis congue. Vivamus metus arcu, adipiscing molestie, hendrerit at, vulputate vitae, nisl.', 533.7);

insert into items (id, image_url, title, description, price) values (11, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/man-1850181_1920_kxidxr.jpg', 'Job', 'Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue. Aliquam erat volutpat.', 337.4);

insert into items (id, image_url, title, description, price) values (12, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/portrait-2194457_1920_rnytv9.jpg', 'Y-Solowarm', 'Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem.
Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy. Integer non velit.
Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque.', 266.0);

insert into items (id, image_url, title, description, price) values (13, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/bicycle-1834265_1920_litf3x.jpg', 'Treeflexr', 'Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque.', 417.0);

insert into items (id, image_url, title, description, price) values (14, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/quickdraws-6154461_1920_qnv6ga.jpg', 'Bitchip', 'Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem.', 635.6);

insert into items (id, image_url, title, description, price) values (15, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/person-690547_1920_u4n1fj.jpg', 'Namfix', 'Sed ante. Vivamus tortor. Duis mattis egestas metus.', 61.8);

insert into items (id, image_url, title, description, price) values (16, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/bottles-774466_1920_d5u48q.jpg', 'Spanr', 'Praesent id massa id nisl venenatis lacinia. Aenean sit amet justo. Morbi ut odio.
Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo. In blandit ultrices enim. Lorem ipsum dolor sit amet, consectetuer adipiscing elit.', 258.0);

insert into items (id, image_url, title, description, price) values (17, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/stew-750846_1920_qdgtea.jpg', 'Home Ing', 'Duis consequat dui nec nisi volutpat eleifend. Donec ut dolor. Morbi vel lectus in quam fringilla rhoncus.
Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci. Mauris lacinia sapien quis libero.
Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh.', 990.8);

insert into items (id, image_url, title, description, price) values (18, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/fly-fishing-1149502_1920_j9819q.jpg', 'Tint', 'Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis.
Sed ante. Vivamus tortor. Duis mattis egestas metus.
Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.', 781.3);

insert into items (id, image_url, title, description, price) values (19, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/hiking-1149891_1920_j0gxw5.jpg', 'Trippledex', 'Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy. Integer non velit.
Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque.', 260.1);

insert into items (id, image_url, title, description, price) values (20, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/tent-548022_1920_kn2iiq.jpg', 'Konklab', 'Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat.', 859.7);

insert into items (id, image_url, title, description, price) values (21, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/rosemary-potatoes-1446677_1920_u6rudl.jpg', 'Treeflex', 'In quis justo. Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet.
Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui.', 411.3);

insert into items (id, image_url, title, description, price) values (22, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840705/man-1850181_1920_kxidxr.jpg', 'Andalax', 'Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque.', 975.3);

insert into items (id, image_url, title, description, price) values (23, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840024/sample.jpg', 'Stringtough', 'Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum.', 158.0);

insert into items (id, image_url, title, description, price) values (24, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/portrait-2194457_1920_rnytv9.jpg', 'Cookley', 'Maecenas tristique, est et tempus semper, est quam pharetra magna, ac consequat metus sapien ut nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris viverra diam vitae quam. Suspendisse potenti.', 367.9);

insert into items (id, image_url, title, description, price) values (25, 'https://res.cloudinary.com/db4plm7gz/image/upload/v1625840706/bicycle-1834265_1920_litf3x.jpg', 'Stronghold', 'Maecenas ut massa quis augue luctus tincidunt. Nulla mollis molestie lorem. Quisque ut erat.
Curabitur gravida nisi at nibh. In hac habitasse platea dictumst. Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem.
Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat.', 318.3);
