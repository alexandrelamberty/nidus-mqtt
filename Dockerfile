FROM golang:1.19-alpine
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /mqtt cmd/main.go
EXPOSE 3333
CMD [ "/mqtt" ]