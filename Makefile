VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`
LDFLAGS = "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILDTIME)"
build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o gosky -ldflags="-s -w -X main.Version=${VERSION} -X main.BuildTime=${BUILDTIME}"  main.go && go clean -cache
# 当前目录执行
vet:
	go vet ./...

test:
	go test ./... -cover

bench:
	go test ./...  -test.bench . -test.benchmem=true

serve:
	go run -ldflags $(LDFLAGS) main.go serve --env=local

job:
	go run -ldflags $(LDFLAGS) main.go job $(filter-out $@, $(MAKECMDGOALS))

wire:
	cd app && wire ./...

%:
	@true


