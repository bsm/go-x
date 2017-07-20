default: test

test:
	go test ./...

bench:
	go test ./... -run=NONE -bench=. -benchmem

README.md: README.md.tpl $(wildcard *.go)
	becca -package $(subst $(GOPATH)/src/,,$(PWD))
