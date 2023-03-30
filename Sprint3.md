## Sprint 3

Work completed in Sprint 3:
  - updated backend structure and database structure to support servers groups and unique users

Frontend tests:
  - Cypress e2e test that begins at the homepage, navigates to account-create, and tries to create an already existing account.

Backend tests:
  - 

Backend API Documentation:
  - Database schema: users,groups,servers(Servers have users and groups, groups have users)
  - three get functions for each entity (getAll, getName, getID) 
  - get all returns all entities in relationship
  - getname performs a search by name nad returns error or matches depending on result
  - getID performs a search by ID nad returns error or matches depending on result
  - each entity has a Delete request
  - for servers you can delete users and groups
  - for group you can delete users
  - each entity has put request to add new tupples to specific tables
  - for servers you can create a server and add users or groups
  - for groups you can create a group and add users
  - for users you can create a user(duplicate usernames not allowed)
