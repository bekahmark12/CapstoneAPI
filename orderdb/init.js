db.createUser(
    {
        user : "admin",
        pwd : "pass",
        roles : [
            {
                role : "readWrite",
                db : "order-db"
            }
        ]
    }
)