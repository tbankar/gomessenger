FROM golang:1.10.4

LABEL version="1.0"

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN mkdir -p /go/src/gomessenger/apig
RUN mkdir -p /go/src/gomessenger/pkg/proto


ADD ./apig /go/src/gomessenger/apig
ADD ./pkg/proto /go/src/gomessenger/pkg/proto
COPY ./Gopkg.toml ./Gopkg.lock /go/src/gomessenger/

WORKDIR /go/src/gomessenger/apig/cmd


RUN dep ensure -v

# Build the service inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apig .
CMD ["./apig"]



