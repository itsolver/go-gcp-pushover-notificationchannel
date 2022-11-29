ARG GOLANG_VERSION="1.18"

ARG PROJECT="go-gcp-pushover-notificationchannel"

ARG COMMIT
ARG VERSION

FROM docker.io/golang:${GOLANG_VERSION} as build

ARG PROJECT
WORKDIR /${PROJECT}

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/server cmd/server
COPY pushover pushover
COPY function.go function.go
COPY incident_type.go incident_type.go
COPY incident_type_test.go incident_type_test.go
COPY message.go message.go
COPY template.go template.go

ARG COMMIT
ARG VERSION
RUN BUILD_TIME=$(date +%s) && \
    CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags "-X 'main.BuildTime=${BUILD_TIME}' -X 'main.GitCommit=${COMMIT}' -X 'main.OSVersion=${VERSION}'" \
    -o /bin/server \
    ./cmd/server


FROM gcr.io/distroless/base-debian11

LABEL org.opencontainers.image.source https://github.com/DazWilkin/go-gcp-pushover-notificationchannel

COPY --from=build /bin/server /

ENV PUSHOVER_USERKEY=""
ENV PUSHOVER_TOKEN=""

ENV PORT="8080"

ENTRYPOINT ["/server"]
