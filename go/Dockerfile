FROM golang:1.24

RUN addgroup -S mercari && adduser -S trainee -G mercari
USER trainee
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of code
COPY app ./app
COPY cmd ./cmd
RUN CGO_ENABLED=1 go build ./cmd/api

# Run our backend
CMD ["./api"]
