FROM golang

RUN mkdir -p /go/src/ebanking-gokit

ADD . /go/src/ebanking-gokit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/haunguyenctnp/ebanking-gokit/transactions/cmd  -watch github.com/haunguyenctnp/ebanking-gokit/transactions
