FROM golang AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

FROM golang as runner
WORKDIR /app

COPY --from=builder /app/serve .

CMD ["./serve"]
