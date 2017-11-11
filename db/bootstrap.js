// Bootstraps the Mongo database.

//
// Create indexes.
//

db.users.createIndex({ "profile.email": 1 }, { unique: true });
