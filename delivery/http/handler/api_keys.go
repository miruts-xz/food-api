package handler

const userApiKey = "9fe8c794-32fd-4dba-b57e-68194327285d"
const adminApiKey = "3ce84e01-b52b-40fb-ab6c-e34643257d4a"

/*

userApiKey = 9fe8c794-32fd-4dba-b57e-68194327285d
adminApiKey = 3ce84e01-b52b-40fb-ab6c-e34643257d4a

food-item request auth:
	getItems: user or admin
	getItem: user or admin
	postItem: admin
	putItem: admin
	deleteItem: admin
order request auth:
	getOrders: user or admin
	getOrder: user or admin
	putOrder: user or admin
	postOrder: user or admin
	deleteOrder: user or admin
user request auth:
	getUsers: admin
	getUser: user or admin
	postUser: user or admin
	putUser: user or admin
	deleteUser: user or admin
	byUsername: user or admin

dart usage:
	headers: <String, String>{
			'api-key': "9fe8c794-32fd-4dba-b57e-68194327285d",
		 	}
*/
