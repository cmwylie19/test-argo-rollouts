FROM golang:latest
WORKDIR /go/src/github.com/cmwylie19/test-argo-rollouts/
RUN go get -d -v golang.org/x/net/html  
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/cmwylie19/test-argo-rollouts/main .
CMD ["./main"]  