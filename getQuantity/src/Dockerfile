FROM golang:latest
ENV TIER=development
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
#RUN go mod download
RUN go build -o main .
EXPOSE 8083
CMD ["./main"]

