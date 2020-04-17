FROM golang:1.14

WORKDIR $GOPATH/src/github.com/naveenchander/GoKube

COPY GoKube.exe .

ENV Port=8000
EXPOSE $Port

CMD GoKube.exe
