FROM golang:1.24

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of code
COPY app ./app
COPY cmd ./cmd
RUN CGO_ENABLED=1 go build ./cmd/api

RUN addgroup --system mercari && adduser --system --ingroup mercari trainee
RUN chown -R trainee:mercari /app
USER trainee

# Run our backend
CMD ["./api"]
