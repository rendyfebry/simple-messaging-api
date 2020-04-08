# Start by building the application.
FROM golang:1.12 as build

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/rendyfebry/simple-messaging-api
COPY . .

RUN rm -rf vendor .vendor*
RUN make build


# Now copy it into our base image.
FROM gcr.io/distroless/base

# Copy bin file
COPY --from=build /go/src/github.com/rendyfebry/simple-messaging-api/bin/msgapi /msgapi

# Copy public folder
COPY public /public

ENTRYPOINT ["/msgapi"]
EXPOSE 8080