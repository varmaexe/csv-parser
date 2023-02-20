# CSV Parser

# Overview
This project is a Golang-based API that parses CSV files and returns the data in JSON format. The API accepts CSV files with a header row, and it assumes that the first row contains the column names. The API uses the encoding/csv package in Golang to read the CSV file, and it returns the data as slice of structs in formatted way.

# Installation

To use the API, you will need to have Golang installed on your computer. You can download and install Golang from the official website: https://golang.org/dl/

Once you have Golang installed, you can clone the repository and run the following command to install the required packages:

```go
go get github.com/gorilla/mux
```

# Usage
To use the API, run the following command:

```go
go run main.go
```

This will start the server on port 8000. You can then make a POST request to the API with a CSV file in the request body. The API will parse the CSV file and return the data in formatted way.

API Endpoints
The API has one endpoint:

POST /upload
This endpoint accepts a CSV file in the request body and returns the data in slice of structs

# Contributing
If you would like to contribute to the project, please follow these steps:

Fork the repository.
1. Create a new branch for your changes.
2. Make your changes and test them.
3. Commit your changes and push them to your fork.
4. Create a pull request.
