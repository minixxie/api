FROM golang:1.8.1-alpine
WORKDIR /go/src/api
RUN apk update && apk add --no-cache git && apk add --no-cache bash
  
RUN go get gopkg.in/yaml.v2
RUN go get gopkg.in/labstack/echo.v3
RUN go get github.com/labstack/echo/middleware
RUN go get gopkg.in/go-on/go.uuid.v1
RUN go get github.com/lib/pq
RUN go get github.com/dgrijalva/jwt-go

ADD . /go/src/api
RUN cd /go/src/api && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=0 /go/src/api/main .
CMD ["/main"] 
