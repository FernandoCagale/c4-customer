# c4-customer

## Flow [c4-kustomize](https://github.com/FernandoCagale/c4-kustomize)

## Dependencies

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

## Build Docker

`build and publish c4-customer`

```sh
$   ./scripts/publish
```

## Kubernetes [YAML](https://github.com/FernandoCagale/c4-kustomize/tree/master/c4-customer/base)

    *   deployment.yaml
    *   service.yaml
    *   virtualservice.yaml

## Running local

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

```sh
$   ./scripts/start.sh
```