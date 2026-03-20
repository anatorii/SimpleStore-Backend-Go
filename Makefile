.PHONY: shop clean

server:
	go build -o shop ./cmd/shop/main.go
	./shop

clean:
	rm -f ./shop
