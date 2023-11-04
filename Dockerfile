FROM golang:1.21.1 AS build
WORKDIR /usr/src/app
COPY . .
RUN ["go", "mod", "tidy"]
RUN ["make"]
EXPOSE 8080
CMD ["./main"]