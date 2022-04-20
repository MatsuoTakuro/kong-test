# setup

## install just

```
brew install just
```

## add .env

```
IP_ADDRESS= `${YOUR_IP_ADDRESS}`
```

## run API server and kong gateway

```
just run
```

# how to use

## request

```
curl -X POST -d '{"name":"hogehoge"}' http://localhost:8000/store-api
```

## response

```
hogehoge
```
