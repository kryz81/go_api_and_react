### Simple PoC for Go REST API with embedded frontend app

It shows an example of embedding a frontend app (here React.js) into compiled Go application

#### Requirements
* Go
* MongoDB
* Node.js

#### Install
* ```go mod tidy```
* ```cd frontend```
* ```npm i```

#### Run
```go run .``` or install **air** and run ```air```

#### Build and run compiled version
* ```go build -o build```
* ```cd build```
* ```APP_PORT=4000 DB_URI=127.0.0.1:27017 DB_NAME=api ./go_api_and_react``` 
