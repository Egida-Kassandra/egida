go-path:
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
	export CGO_ENABLED=0

build: go-path
	go build -o build/egida cmd/main.go


docs:
	mkdocs gh-deploy