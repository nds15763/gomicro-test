build:
    protoc -I. --go_out=plugins=grpc:. proto/demo/demo.proto
    GOOS=linux GOARCH=amd64 go build
    docker build -t laracom-demo-service .
run:
    docker run -d -p 9091:9999 -p 8081:8000 laracom-demo-service