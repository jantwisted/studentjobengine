# Build Stage
FROM golang:latest

RUN apt-get update
RUN apt-get install make
RUN apt-get install git

RUN curl https://glide.sh/get |sh

LABEL app="build-studentjobengine"
LABEL REPO="https://github.com/jantwisted/studentjobengine"

ENV PROJPATH=/go/src/github.com/jantwisted/studentjobengine

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/jantwisted/studentjobengine
WORKDIR /go/src/github.com/jantwisted/studentjobengine
RUN glide install

RUN make build
RUN mkdir -p /opt/studentjobengine/bin
RUN cp ./bin/studentjobengine /opt/studentjobengine/bin

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/studentjobengine/bin

WORKDIR /opt/studentjobengine/bin

RUN chmod +x /opt/studentjobengine/bin/studentjobengine

CMD ["/opt/studentjobengine/bin/studentjobengine"]
