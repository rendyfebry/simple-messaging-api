# Simple Messaging API

Simple Messaging API with REST & Websocket

## Prerequisite

- Golang 1.9
- Dep

## Install depedencies

```
dep ensure
```

## Test

```
go test -v ./...
```

## Usage

Run Server

```
go run main.go
```

### Webclient

Visit the url on browser

```
http://localhost:8080/
```

PS: This webclient implement the `/ws` websocket service, to post and listed message in realtime.

### API

CreateNewMessage

```
curl -L -X GET 'localhost:8080/messages'
```

GetMessages

```
curl -L -X POST 'localhost:8080/messages' \
	-H 'Content-Type: application/json' \
	--data-raw '{
		"body": "Hai!"
	}'
```

PS: The api and webclient is connected, so weather post message from api or web client, the message will appear when you call the GetMessages endpoint.

## Docker

Docker Build

```
docker build -t msgapi .
```

Run in Docker

```
docker run --rm --name msgapi \
        -p 8080:8080 \
        -e "USR_ENV=development" \
        msgapi
```
