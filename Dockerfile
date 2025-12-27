FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 


# copy the source code from the build context to the workdir 
COPY . .

# Build the go app for linux without any dependency for the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp ./cmd/api 


FROM alpine:latest

WORKDIR /app 


#copying  the binary from the  builder stage
COPY --from=builder /app/myapp ./

CMD ["./myapp"]


