FROM golang:1.14 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /workspace

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./build/websay .

FROM gcr.io/distroless/base
COPY --from=builder /workspace/build/websay /websay

ENTRYPOINT [ "/websay" ]