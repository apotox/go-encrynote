var rolesToGrant = ["dbAdmin"];
print("Granting roles ++++++++++++++++++++++++++++++++++++ ");
var rootDb = db;
var admin = db.getSiblingDB("admin");
var users = admin.system.users.find({}).toArray();

users.forEach(({ user, db }) => {
    print(`${db}: Granted ${user} the roles ${JSON.stringify(rolesToGrant)}.`);
    rootDb.getSiblingDB(db).grantRolesToUser(user, rolesToGrant);
});
