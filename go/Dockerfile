FROM golang:1.24 AS build

WORKDIR /app

# Install dependencies (layer 1)
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of code (layer 2)
COPY app ./app
COPY cmd ./cmd
RUN CGO_ENABLED=0 go build ./cmd/api

# Multi-stage build
FROM scratch

COPY --from=build /app/api .

# Run our backend
CMD ["./api"]
