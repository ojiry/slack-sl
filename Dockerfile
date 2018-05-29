FROM golang

ADD . /go/src/github.com/ojiry/slack-sl
RUN go install github.com/ojiry/slack-sl/cmd/sl

CMD ["/go/bin/sl"]
