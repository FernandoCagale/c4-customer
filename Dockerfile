## Builder Image
FROM golang:1.13-stretch as builder
ENV GO111MODULE=on
COPY . /c4-customer
WORKDIR /c4-customer
RUN go get github.com/google/wire/cmd/wire
RUN wire ./...
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application cmd/c4/main.go cmd/c4/wire_gen.go

## Run Image
FROM scratch
COPY --from=builder /c4-customer/cmd/c4/bin/application application
EXPOSE 8080
ENTRYPOINT ["./application"]