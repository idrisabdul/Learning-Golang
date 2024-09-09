# Latest golang image on apline linux
FROM golang:1.22.1-alpine

ENV TZ="Asia/Jakarta"

RUN apk add tzdata
RUN ln -s /usr/share/zoneinfo/Asia/Jakarta /etc/localtime


# Work directory
WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Starting our application
CMD ["go", "run", "cmd/main.go"]

# Exposing server port
EXPOSE 8080

