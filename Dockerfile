# Build stage
FROM golang AS build-env

WORKDIR /go/src/app
COPY * ./

RUN go get -v -d
RUN go install -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o gocave .

# final stage
FROM scratch

COPY --from=build-env /go/src/app/gocave /

ENTRYPOINT [ "/gocave" ]
