FROM arukasio/arukas:dev
MAINTAINER "Shuji Yamada <s-yamada@arukas.io>"
COPY . $GOPATH/src/github.com/arukasio/cli
WORKDIR $GOPATH/src/github.com/arukasio/cli
RUN godep restore -v .../.
RUN ARUKAS_DEV=1 scripts/build.sh
WORKDIR $GOPATH
RUN touch .env
ENTRYPOINT ["bin/arukas"]
