FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o main .

RUN apk add curl
RUN curl -sSL https://raw.githubusercontent.com/cyberark/summon/main/install.sh | sh
RUN curl -sSL https://raw.githubusercontent.com/cyberark/summon-conjur/main/install.sh | bash

CMD ["/usr/local/bin/summon","-f","/app/secrets.yml","-p","/app/summon-conjur","/app/main"]
