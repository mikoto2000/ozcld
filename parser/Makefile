all: yacc build run

run: cldparser
	./cldparser testdata/target.ozcld > /tmp/ozcld.dot
	dot -Tpng /tmp/ozcld.dot -o /tmp/ozcld.png
	img2sixel /tmp/ozcld.png
	rm /tmp/ozcld.png
	rm /tmp/ozcld.dot

build: cldparser.go cldparser.go
	go build -o cldparser cldparser.go

yacc: cldparser.go.y
	go tool yacc -o cldparser.go cldparser.go.y

.PHONY:clean
clean:
	rm cldparser.go
	rm *.output
	rm cldparser
