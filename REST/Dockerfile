FROM golang

WORKDIR /go/src

RUN go mod init src

RUN go get -u github.com/gorilla/mux

RUN go get github.com/go-sql-driver/mysql

COPY . .

CMD ["go", "run", "."]

