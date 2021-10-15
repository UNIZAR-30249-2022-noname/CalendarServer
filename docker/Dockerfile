FROM golang:alpine

MAINTAINER arejula27 


# Choose te directory where all action will happen
WORKDIR /app

ENV GIN_MODE=release
ENV PORT=8080

#Copy the whole project inside the container
COPY . .

#Get git
RUN apk update && apk add --no-cache git

#Download dependencies
RUN go get ./...

#Compile
RUN go build ./src/internal/cmd/

EXPOSE $PORT

ENTRYPOINT ["./cmd"]


