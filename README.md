# Go Clean Architecture

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

