# Dependencies and Binary Builder
FROM golang:1.24-alpine AS builder

WORKDIR /usr/src/app/

# Install build tools (git, C compiler) and libopus headers
RUN apk update && apk add --no-cache build-base git opus-dev

# Explicitly enable CGO (important for gopus)
ENV CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download \
    && go install github.com/bwmarrin/dca/cmd/dca@latest

COPY . .
RUN go build -v -o /yona main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

RUN apk update \
    && apk add --no-cache ffmpeg opus

# Copy assets dir
COPY --from=builder /usr/src/app/assets/ /app/assets/
# Copy dca bin
COPY --from=builder /go/bin/dca /usr/bin/dca
# Copy yona bin
COPY --from=builder /yona /usr/bin/yona


CMD ["yona"]