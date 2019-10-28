FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/ojiry/slack-sl
COPY . .
RUN go build main.go

FROM alpine
COPY --from=builder /go/src/github.com/ojiry/slack-sl /app

CMD /app/main
