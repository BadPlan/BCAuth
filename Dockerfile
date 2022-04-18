FROM golang:1.18
RUN mkdir /app
COPY . /app
WORKDIR /app/cmd
RUN mkdir /var/log/bcauth
RUN chmod 0777 /var/log/bcauth
EXPOSE 8090
RUN go build -o BCAuth .
CMD ["/app/cmd/BCAuth"]
