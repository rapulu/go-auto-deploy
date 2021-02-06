FROM golang:1.14-alpine

ADD . /home
        
WORKDIR /home

CMD ["go","run","main.go"]