APP?=GoKube.exe
PORT?=8000
RELEASE?=0.0.2

clean:
		del /f ${APP}

build: clean
		go build -ldflags="-s -w -X main.Release=${RELEASE} -X main.PORT=${PORT}" -o ${APP}

run: build
		./${APP}

test:
		go test -v -race ./...