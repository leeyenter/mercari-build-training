FROM golang:1.24

WORKDIR /app
RUN addgroup --system mercari && adduser --system --ingroup mercari trainee

# Install dependencies (layer 1)
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of code (layer 2)
COPY app ./app
COPY cmd ./cmd
RUN CGO_ENABLED=1 go build ./cmd/api

# Configure user (layer 3)
RUN chown -R trainee:mercari /app
USER trainee

# Run our backend
CMD ["./api"]
