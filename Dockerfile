FROM golang:1.12.9-stretch@sha256:44600a24dff9a70122d4446f63903a642e81c0422cd0d87249a8a5183ba5f926

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

RUN go test

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags '-w -extldflags "-static"'

#FROM scratch
FROM busybox@sha256:9f1003c480699be56815db0f8146ad2e22efea85129b5b5983d0e0fb52d9ab70

COPY --from=0 /src/gotemplate /usr/local/bin/gotemplate

ENTRYPOINT ["/usr/local/bin/gotemplate"]
