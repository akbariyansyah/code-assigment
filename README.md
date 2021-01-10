### Code Assignment

## How to run this project

### Without Docker
1. configure your config.json
2. import database.sql
3. run the project

*  build project

```
$ go build -o main
```
* Run app
```
$ ./main
```
* Run Test
```
$ go test -v ./...
```

### With docker

```
$ docker-compose -f docker-compose.yml up -d
```

### Using Make

```
$ make run
```
## API Spec

### 1. Check Saldo

- Request :
    - Method : GET
    - Endpoint : ```/account/{account_number}```
    - Header :
       - Accept : Application/json

- Response :
    - Code : 200
 ```json
{
    "account_number" : "555001",
    "customer_name" : "Bob Martin",
    "balance" : 10000
}

```

### 2. Transfer

- Request :
    - Method : POST
    - Endpoint ```/account/{from_account_number}/transfer```
    - Header :
         - Accept : Application/json
         - Content-Type: application/json
    - Body :
 ```json
  {
    "to_account_number" : "555002",
    "amount" : 100
  }
  ```
- Response :
    - Code : 201     
