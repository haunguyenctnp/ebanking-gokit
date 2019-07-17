FROM golang

RUN mkdir -p /go/src/github.com/haunguyenctnp/ebanking-gokit

ADD . /go/src/github.com/haunguyenctnp/ebanking-gokit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/haunguyenctnp/ebanking-gokit/authentication/cmd  -watch github.com/haunguyenctnp/ebanking-gokit/authentication
