FROM golang:1.10.4

LABEL version="1.0"

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN cd /opt/gomessenger/cmd && dep ensure -v

WORKDIR /opt/gomessenger/cmd

# Build the service inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apig .
CMD ["./apig"]



