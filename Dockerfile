FROM golang:1.9 as builder

RUN curl https://glide.sh/get | sh

WORKDIR  /go/src/github.com/kontena/pharos-default-backend

# Add dependency graph and vendor it in
ADD glide.yaml glide.lock /go/src/github.com/kontena/pharos-default-backend/
RUN glide install

# Add source and compile
ADD *.go /go/src/github.com/kontena/pharos-default-backend/

ARG arch=amd64
ARG arm_level=7

# Compile based on the build args
RUN \
    if [ "$arch" = "arm64" ]; then \
        CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=${arm_level} go build -a -installsuffix cgo -ldflags '-w' -o server . ;\
    else \
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o server . ;\
    fi

FROM scratch

COPY --from=builder /go/src/github.com/kontena/pharos-default-backend/server .

ADD static /static/

CMD ["./server"]