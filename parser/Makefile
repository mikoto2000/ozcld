all: yacc build run

run: cldparser
	./cldparser test04.txt

build: cldparser.go cldparser.go
	go build -o cldparser cldparser.go

yacc: cldparser.go.y
	go tool yacc -o cldparser.go cldparser.go.y

.PHONY:clean
clean:
	rm cldparser.go
	rm *.output
	rm cldparser
