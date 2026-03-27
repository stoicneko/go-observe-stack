FROM golang
WORKDIR /app
COPY main.go .
RUN go mod init myapp
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o server main.go
CMD ["/app/server"]
