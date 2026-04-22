FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o my-go-app main.go
EXPOSE 8080
CMD ["./my-go-app"]