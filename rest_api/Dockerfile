# iron/go:dev is the alpine image with the go tools added
FROM golang:1.11.2-alpine 
EXPOSE 8000
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/rest_api
ADD . $GOPATH/src/rest_api

WORKDIR $GOPATH/src/rest_api
RUN go build -o rest_api . 
CMD ["/go/src/rest_api/rest_api"]
