all: build

build: 
	go build -o ./bin/portfolio ./cmd/portfolio/main.go 

start: 
	./cmd/portfolio/main.go 

dev: 
	go run cmd/portfolio/main.go

watch-css: 
	npx tailwindcss -o static/styles.css --watch

build-css: 
	npx tailwindcss -o static/styles.css
