FROM golang:1.18.0-alpine3.15 as builder
RUN apk --no-cache add git dep ca-certificates
ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"
RUN mkdir -p $GOPATH/app/github.com/figueyes/shortener-app
WORKDIR $GOPATH/app/github.com/figueyes/shortener-app
COPY go.mod .
COPY go.sum .
COPY app/ app/
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a \
  -installsuffix cgo -ldflags '-extldflags "-static"' \
  -o $GOBIN/main ./app/main.go

FROM scratch as production
ARG VERSION
ENV VERSION_APP=$VERSION
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/ /app/
EXPOSE 3000
ENTRYPOINT ["/app/main"]