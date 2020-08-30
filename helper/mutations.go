package main

const createUserMutation = "mutation CreateUser($input: [AddUserInput!]!) { addUser(input: $input) { user { email } } }"

const editUserMutation = "mutation EditUser($input: UpdateUserInput!) { updateUser(input: $input) { user { email } } }"

const removeUserMutation = "mutation RemoveUser($filter: UserFilter!) { deleteUser(filter: $filter) { user { email } } }"
