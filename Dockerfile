FROM golang:1.16-alpine as builder

WORKDIR /app
COPY . .
RUN go get
RUN go build -o server_demo server_demo.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server_demo .
CMD ["./server_demo"]
# RUN go get -d -v ./...
# RUN go install -v ./...

# ENV GIN_MODE=release

# CMD ["orotig_api"]

# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

# RUN go build -o /server_demo

# EXPOSE 8080

# CMD [ "/server_demo" ]