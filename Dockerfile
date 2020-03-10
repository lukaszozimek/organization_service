FROM golang

RUN mkdir -p /go/src/github.com/lukaszozimek/organization_service

ADD . /go/src/github.com/lukaszozimek/organization_service

RUN cd /go/src/github.com/lukaszozimek/organization_service && go mod download

ENTRYPOINT  watcher -run github.com/lukaszozimek/organization_service/cmd -watch github.com/lukaszozimek/organization_service
