.PHONY: run build clean

# Go uygulamasını çalıştır
run:
	go run cmd/server/main.go

# Go uygulamasını derle
build:
	go build -o bin/server cmd/server/main.go

# Derlenen dosyaları temizle
clean:
	rm -rf bin/

# Varsayılan hedef
default: run
