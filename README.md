# HNGStage2

## API Endpoints

- **POST /api**: Creates a user.
- **GET /api/{userID}**: Retrieve details of a specific person.
- **PUT /api/{userID}**: Update an existing person.
- **DELETE /api/{userID}**: Delete a specific person.

## GET localhost:10000/api/{userID}
```
[localhost:10000/api/8
```

Description: Retrieves a specific person..

Method: GET

Returns: A specific user.

## POST localhost:10000/api/
```
[localhost:10000/api/
```

Description: creates a new user.

Method: POST

Request Body: Contains the necessary data to create a user.
```
{
   "Name" : "Adebayo David"
}
```

## PUT localhost:10000/api/{userID}
```
[localhost:10000/api/8
```

Description: Updates the details of a person.

Method: PUT

Request Body: Contains the updated information for the person.
```
{
   "Name" : "Adebayo David"
}
```

## DELETE localhost:10000/api/{userID}
```
localhost:10000/api/8
```

Description: Deletes a specific user.

Method: DELETE


## How to run this project

- **RUN "go get" FOR EACH OF THE FOLLOWING ON YOUR TERMINAL
  
   github.com/go-sql-driver/mysql v1.5.0 // indirect
  
	github.com/gorilla/mux v1.8.0 // indirect

	github.com/jinzhu/gorm v1.9.16 // indirect

	github.com/jinzhu/inflection v1.0.0 // indirect

	github.com/joho/godotenv v1.5.1 // indirect

	github.com/lib/pq v1.10.9 // indirect

- **RUN "go run main.go"


