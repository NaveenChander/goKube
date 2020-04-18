APP?=GoKube.exe
DOCKERTAG?=gokube
DOCKERPORT?=8000
HOSTPORT?=8000
RELEASE?=0.0.2

clean:
		del /f ${APP}

build: clean
		go build -ldflags="-s -w -X main.Release=${RELEASE} -X main.PORT=${DOCKERPORT}" -o ${APP}

container: build
		docker build -t $(DOCKERTAG):$(RELEASE) .

run: container		
		docker stop $(DOCKERTAG) || true
		docker rm $(DOCKERTAG) || true
		docker run -p ${HOSTPORT}:${DOCKERPORT} --name $(DOCKERTAG) $(DOCKERTAG):$(RELEASE)

#### to Run the app without a docker		./${APP}

test:
		go test -v -race ./...