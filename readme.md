# Max Integrations Admin API
maxinteg-admin-go is a Go project that provides a server API and basic CRUD operations for a project management application.

## Requirements

- Go 1.14 or higher
- [Gorilla Mux](https://github.com/gorilla/mux)

## Installation

1. Clone the repo

`git clone https://github.com/maxinteg/maxinteg-admin-go.git`

2. Install dependencies

`go get -u -v ./...`

3. Run the app

`go run main.go`

## Usage

The API provides the following endpoints for basic CRUD operations for a shopping list application:

- POST /user - Create a new user
- GET /user/{id} - Retrieve a user by ID
- PUT /user/{id} - Update a user by ID
- DELETE /user/{id} - Delete a user by ID
- POST /item - Create a new shopping item
- GET /item/{id} - Retrieve a shopping item by ID
- PUT /item

# db

This is a Go package for handling user authentication with Firebase. It includes a function that checks a user's credentials against Firebase and returns a JSON response with the user's data if the credentials are valid. It also uses the bcrypt library to securely hash and compare user passwords.