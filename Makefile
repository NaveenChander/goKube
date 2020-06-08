APP?=GoKube.exe
DOCKERTAG?=gokube
DOCKERPORT?=8000
HOSTPORT?=8000
RELEASE?=0.0.2
DBSERVER?=localhost
DBPORT?=1433
DBUSER?=sa
DBPASSWORD?=abc@123
DBCATALOGUE?=GoKube
EXPCachePeriod?=1440

clean:
		del /f ${APP}

build: clean
		go build -ldflags="-s -w \
		-X configuration.Release=${RELEASE} \
		-X configuration.PORT=${DOCKERPORT} \
		-X configuration.DBSERVER=${DBSERVER} \
		-X configuration.DBPORT=${DBPORT} \
		-X configuration.DBUSER=${DBUSER} \
		-X configuration.DBPASSWORD=${DBPASSWORD} \
		-X configuration.DBCATALOGUE=${DBCATALOGUE} \
		-X configuration.EXPCachePeriod=${EXPCachePeriod}" \
		-o ${APP}

container: build
		docker build -t $(DOCKERTAG):$(RELEASE) .

run: container		
		docker stop $(DOCKERTAG) || true
		docker rm $(DOCKERTAG) || true
		docker run -p ${HOSTPORT}:${DOCKERPORT} --name $(DOCKERTAG) $(DOCKERTAG):$(RELEASE)

#### to Run the app without a docker		./${APP}

test:
		go test -v -race ./...