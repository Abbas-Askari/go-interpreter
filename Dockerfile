FROM golang:1.24-alpine AS builder


WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: Final image
FROM scratch

COPY --from=builder /app/main /main

CMD ["/main"]