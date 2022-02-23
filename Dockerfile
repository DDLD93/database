FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/ddld93/database/
RUN go get -d -v golang.org/x/net/html  
COPY app.go    ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o formApp .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/ddld93/auth/formApp ./
RUN mkdir images
CMD ["./formApp"]  
