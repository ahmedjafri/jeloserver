# For Go 1.10
FROM golang:1.10

WORKDIR $GOPATH/src/github.com/ahmedjafri/jeloserver
COPY . .

RUN go build -o /app/main . 
CMD ["/app/main"]

EXPOSE 3000