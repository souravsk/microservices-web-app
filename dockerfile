#This first phase is just to build the go we are not using it to exeute becasue it take to much space
FROM golang:latest as builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main .
# EXPOSE 80
# ENTRYPOINT ["./main"]

#This second phase will we will execute the main file that we build and we use a lite go
FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main .
EXPOSE 80
CMD ["/main"]