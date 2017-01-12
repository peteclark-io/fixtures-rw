FROM golang:1.8-alpine

COPY . /source/

RUN apk add --update bash \
  && ls -lta /source/ \
  && apk --update add git go ca-certificates \
  && cd /source/ \
  && export GOPATH=/gopath \
  && echo "{\"version\": \"$(git describe --tag --always 2> /dev/null)\", \"built\": \"$(date -u +%Y%m%d%H%M%S)\", \"repository\": \"$(git config --get remote.origin.url)\", \"revision\": \"$(git rev-parse HEAD)\", \"go\": \"$(go version)\"}" > /version.json \
  && REPO_PATH="github.com/peteclark-io/match-rw/" \
  && mkdir -p $GOPATH/src/${REPO_PATH} \
  && cp -r /source/* $GOPATH/src/${REPO_PATH} \
  && cd $GOPATH/src/${REPO_PATH} \
  && go get ./... \
  && go build \
  && mv ./match-rw / \
  && apk del go git \
  && rm -rf $GOPATH /var/cache/apk/*

EXPOSE 80

CMD [ "/match-rw" ]
