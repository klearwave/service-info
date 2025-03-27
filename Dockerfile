#
# build image
#
FROM golang:1.23-alpine AS build
WORKDIR /src

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY pkg/ pkg/
COPY migrations/ migrations/

RUN CGO_ENABLED=0 go build -o /src/service ./pkg/cmd && \
        chmod +x /src/service
RUN export GOBIN=/src && go install github.com/pressly/goose/v3/cmd/goose@v3.24.1 && \
        chmod +x /src/goose

#
# runtime image
#
FROM gcr.io/distroless/static:nonroot AS runtime
USER 65532:65532
COPY --from=build /src/migrations /migrations
COPY --from=build /src/service /bin/service
COPY --from=build /src/goose /bin/goose
ENTRYPOINT ["service"]
