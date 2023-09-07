FROM golang:1.20-alpine 


# Set the working directory inside the container
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./
 

RUN go build -buildvcs=false

EXPOSE 3000

CMD [ "./go-microservice" ]