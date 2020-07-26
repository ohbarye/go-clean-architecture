# Go Clean Architecture

This is an example application to implement an API server with Clean Architecture.

## Architecture

All dependencies are unidirectional by using DIP.

For example, unless the `SqlHandler` interface doesn't exist, `UserRepository` directly depends on `SqlHandler` across database layer and infrastructure layer. The same goes on `UserRepository` and `UserInteractor`. 

![image](https://user-images.githubusercontent.com/1811616/88475987-ada7db80-cf6f-11ea-82b5-efd0b46d5423.png)

## Stacks

- MySQL 8.0
- Go 1.14
  - Echo

## Development

```shell
$ docker-compose up
```

```shell
// User create
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"FirstName": "Kentaro", "LastName": "Yoshioka"}' http://localhost:1313/users
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Date: Sun, 26 Jul 2020 08:33:05 GMT
Content-Length: 53

{"ID":2,"FirstName":"Kentaro","LastName":"Yoshioka"}


// User index
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" http://localhost:1313/users
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sun, 26 Jul 2020 08:33:39 GMT
Content-Length: 104

[
  {"ID":1,"FirstName":"Susan","LastName":"Taylor"},
  {"ID":2,"FirstName":"Kentaro","LastName":"Yoshioka"}
]

// User show
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" http://localhost:1313/users/2
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sun, 26 Jul 2020 08:34:12 GMT
Content-Length: 53

{"ID":2,"FirstName":"Kentaro","LastName":"Yoshioka"}
```

### Connect to MySQL

```shell
$ mysql -u root -h localhost -P 3306 --protocol=tcp
```

## Original Idea

- https://github.com/hirotakan/go-cleanarchitecture-sample
  - https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
