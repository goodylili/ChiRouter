# ChiRouter
Code for MatterMost Article on Building APIs in Go using the Chi Router


Consuming the API
This API allows you to create, read, update, and delete students. Here are some examples of how to use curl to interact with the API endpoints:

Creating a student
To create a student, send a POST request to the / endpoint with a JSON body containing the student's name and study field:

Copy code
curl -X POST \
  http://localhost:8080/ \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "Alice",
	"study": "Computer Science"
}'
Reading a student
To read a student, send a GET request to the /get/{name} endpoint, where {name} is the name of the student you want to read:

Copy code
curl -X GET http://localhost:8080/get/Alice
Updating a student
To update a student, send a PUT request to the /update/{name} endpoint with a JSON body containing the updated student information. The {name} parameter in the endpoint should be the name of the student you want to update:

Copy code
curl -X PUT \
  http://localhost:8080/update/Alice \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "Alice",
	"study": "Mathematics"
}'
Deleting a student
To delete a student, send a DELETE request to the /delete/{name} endpoint, where {name} is the name of the student you want to delete:

Copy code
curl -X DELETE http://localhost:8080/delete/Alice
