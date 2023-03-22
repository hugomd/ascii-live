FROM golang:1.13-alpine AS build-env
ENV GO111MODULE=on
WORKDIR /go/src/github.com/hugomd/ascii-live/
COPY . /go/src/github.com/hugomd/ascii-live/
RUN cd /go/src/github.com/hugomd/ascii-live && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=build-env /go/src/github.com/hugomd/ascii-live/main /
CMD ["/main"]
