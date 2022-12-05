FROM golang:alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
#COPY .env .env
ENV SVC_PORT 8002
EXPOSE 8002
CMD ["./main"]