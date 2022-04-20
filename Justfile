set dotenv-load

run:
    docker build --tag kong-dbless .
    docker run -d --name kong-dbless -p 8000:8000 -p 8001:8001 -p 8002:8002 -p 8443:8443 --add-host=local_dev:$IP_ADDRESS kong-dbless
    go run store/main.go
