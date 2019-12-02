**c4-customer - docker**

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

`Docker build c4-customer`

```sh
$   docker build -t c4-customer .
```

`Docker c4-customer`

```sh
$   docker run -d --name c4-customer -p 8080:8080 c4-customer
```

**c4-customer - local**


```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download wire "dependency injection"`

```sh
$   go get -u github.com/google/wire/cmd/wire
```

`generate wire_gen.go`

```sh
$   wire
```

`generate build`

```sh
$   go build -o bin/application
```


```sh
$   ./bin/application
```