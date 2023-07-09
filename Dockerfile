# Stage 1: Build the Go application
FROM golang:alpine as build

RUN apk update && apk add --no-cache git
WORKDIR /src
COPY . .
RUN go mod tidy
RUN go build -o livecode-book

# Stage 2: Create a minimal image for the app
FROM alpine
WORKDIR /app
COPY --from=build /src/livecode-book /app/
ENTRYPOINT ["/app/livecode-book"]
