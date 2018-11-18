GO = go
DOCKER = docker
IMAGE = go-alpine-sub
CFLAGS = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
CGOBUILD = $(CFLAGS) $(GO) build -o main *.go
#PORT = -p 8080:8080 #$(PORT)

default:
	$(DOCKER) run -it --rm $(IMAGE)
go-build:
	$(GO) build -o main .
cgo-build:
	$(CGOBUILD)
docker-build: cgo-build
	$(DOCKER) build -t $(IMAGE) -f Dockerfile .
clean:
	rm main
