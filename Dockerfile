FROM golang:1.16-alpine

COPY . /app
WORKDIR /app

RUN GO BUILD GoKube

COPY GoKube .

ENV Port=8000
EXPOSE $Port

CMD ["GoKube"]
