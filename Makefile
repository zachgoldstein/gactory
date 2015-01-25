FILES=./*.go

fmt:
	go fmt ${FILES}

deps:
	go get github.com/smartystreets/goconvey

test:
	go test ${FILES}

live-test:
	goconvey

doc:
	pkill godoc; godoc -http=":7080" &