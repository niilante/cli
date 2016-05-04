FROM arukasio/arukas:dev
MAINTAINER "Shuji Yamada <s-yamada@arukas.io>"

COPY . $GOPATH/src/github.com/arukasio/cli
WORKDIR $GOPATH/src/github.com/arukasio/cli

RUN godep restore
RUN go list ./...
RUN for package in $(go list ./...| grep -v vendor); do golint ${package}; done
RUN ARUKAS_DEV=1 scripts/build.sh

WORKDIR $GOPATH

ENTRYPOINT ["bin/arukas"]
