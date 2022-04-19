#Shortener api
By @figueyes

## Architecture

General Diagram

<img src="./assets/shortener-app.architecture.png" alt="creation"/>

## Config
First, you need to configure an .env file with example.env variables 
### Dependencies
#### mongo
- configure a local database with localhost:[PORT] database and collection=shorts

#### redis
- configure a local database with localhost:[PORT] and db:0 
  1. cache database
  2. pub/sub to remove from cache
#### kafka
- configure kafka broker to publish message about observability
- topics:
  1. observer of data
#### zookeeper
- kafka dependence

## Running 
```
go run src/main.go
```

