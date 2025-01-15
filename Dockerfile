#
# build stage
#
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o service
RUN export GOBIN=/src && go install github.com/pressly/goose/v3/cmd/goose@latest

#
# service image
#
FROM alpine AS service
RUN apk add --no-cache ca-certificates
COPY --from=build /src/migrations /migrations
COPY --from=build /src/service /bin/service
COPY --from=build /src/goose /bin/goose
ENTRYPOINT ["service"]
