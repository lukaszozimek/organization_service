FROM golang

RUN mkdir -p /go/src/github.com/lukaszozimek/organization_service

ADD . /go/src/github.com/lukaszozimek/organization_service

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/github.com/lukaszozimek/organization_service && go mod download

ENTRYPOINT  watcher -run github.com/lukaszozimek/organization_service/cmd -watch github.com/lukaszozimek/organization_service
