all: build

build: 
	go build -o ./bin/portfolio ./cmd/portfolio/main.go 

run: 
	go run cmd/portfolio/main.go

watch-css: 
	npx tailwindcss -o static/styles.css --watch

build-css: 
	npx tailwindcss -o static/styles.css
