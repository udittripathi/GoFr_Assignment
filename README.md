# GoFr_Assignment

##To run:
```
go run main.go
```


###This Go program creates a simple HTTP API for managing information about cars. The program uses the gofr framework for handling HTTP requests and interactions with a database. Here's a brief overview of the functionality:

1. Create Car Information:
+ Endpoint: POST /carinfo
+ Parses a JSON request body containing car information.
+ Inserts the car information into a database table named "cars."

2. List All Cars:
+ Endpoint: GET /car
+ Retrieves a list of all cars from the "cars" table in the database.

3. Get Car by ID:
+ Endpoint: GET /car/{id}
+ Retrieves information for a specific car based on the provided car ID.

4. Complete Car Repair:
+ Endpoint: POST /completeRepair/{id}
+ Updates the repair status of a specific car to 'Completed' based on the provided car ID.

5. Delete Car Entry:
+ Endpoint: DELETE /deleteCar/{id}
+ Deletes a car entry from the database based on the provided car ID.

###The program initializes a gofr object, sets up various HTTP endpoints for handling CRUD operations on car information, connects to a database, and starts an HTTP server to listen for incoming requests. The database operations involve inserting, querying, updating, and deleting records in the "cars" table.