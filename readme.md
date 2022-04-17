#Shortener api
By @figueyes

## Architecture

Shortener algorithms
sha256


## Config
First, you need to configure an .env file with example.env variables 
### Dependencies
#### mongo
- host: localhost

#### redis
- addr: localhost 
- db: 0
#### kafka
- broker: localhost
- topics:
  1. observer of data
  2. remover of cached data
#### zookeeper
- kafka dependence

## Running 
```
go run src/main.go
```

