FROM golang:1.10.4

LABEL version="1.0"

RUN mkdir -p /go/src/gomessenger/cmd
RUN mkdir -p /go/src/gomessenger/internal
RUN mkdir -p /go/src/gomessenger/pkg

ADD ./cmd /go/src/gomessenger/cmd
ADD ./internal /go/src/gomessenger/internal
ADD ./pkg /go/src/gomessenger/pkg

WORKDIR /go/src/gomessenger/cmd/msngr
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
COPY ./Gopkg.toml ./Gopkg.lock /go/src/gomessenger/

RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .
CMD ["./msngr"]





