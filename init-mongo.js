db.createUser({
    user: 'root',
    pwd: 'toor',
    roles: [
        {
            role: 'readWrite',
            db: 'alertDB',
        },
    ],
});

// Select DB and create new collection
db = new Mongo().getDB("alertDB");
db.createCollection('users', { capped: false });