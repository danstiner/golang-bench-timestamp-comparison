# Build Stage
FROM lacion/alpine-golang-buildimage:1.11 AS build-stage

LABEL app="build-golang-bench-compare"
LABEL REPO="https://github.com/danstiner/golang-bench-compare"

ENV PROJPATH=/go/src/github.com/danstiner/golang-bench-compare

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/danstiner/golang-bench-compare
WORKDIR /go/src/github.com/danstiner/golang-bench-compare

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/danstiner/golang-bench-compare"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/golang-bench-compare/bin

WORKDIR /opt/golang-bench-compare/bin

COPY --from=build-stage /go/src/github.com/danstiner/golang-bench-compare/bin/golang-bench-compare /opt/golang-bench-compare/bin/
RUN chmod +x /opt/golang-bench-compare/bin/golang-bench-compare

# Create appuser
RUN adduser -D -g '' golang-bench-compare
USER golang-bench-compare

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/golang-bench-compare/bin/golang-bench-compare"]
