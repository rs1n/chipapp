// Bootstraps the Mongo database.

//
// Create indexes.
//

db.users.createIndex({ login: 1 }, { unique: true });

//
// Seed the db.
//
