# Build Stage
FROM debian:latest:1.11 AS build-stage

LABEL app="build-studentjobengine"
LABEL REPO="https://github.com/jantwisted/studentjobengine"

ENV PROJPATH=/go/src/github.com/jantwisted/studentjobengine

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/jantwisted/studentjobengine
WORKDIR /go/src/github.com/jantwisted/studentjobengine

RUN make build-alpine

# Final Stage
FROM jantwisted/studentjobs-api:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/jantwisted/studentjobengine"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/studentjobengine/bin

WORKDIR /opt/studentjobengine/bin

COPY --from=build-stage /go/src/github.com/jantwisted/studentjobengine/bin/studentjobengine /opt/studentjobengine/bin/
RUN chmod +x /opt/studentjobengine/bin/studentjobengine

# Create appuser
RUN adduser -D -g '' studentjobengine
USER studentjobengine

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/studentjobengine/bin/studentjobengine"]
