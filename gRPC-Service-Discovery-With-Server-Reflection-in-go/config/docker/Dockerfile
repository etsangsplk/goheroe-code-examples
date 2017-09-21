FROM golang:alpine
RUN apk update \
  && apk add git
COPY $GOPATH/bin/server $GOPATH/bin/server 
EXPOSE 6666
ENTRYPOINT ["/go/bin/server"]
