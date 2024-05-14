name := eqasciiparser
eq_path := /src/eq/takp/

run: build
	@cd $(eq_path) && wine $(name).exe

build:
	@mkdir -p bin
	@CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o $(eq_path)/$(name).exe *.go