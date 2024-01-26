all: build

build: 
	go build -o ./bin/portfolio ./cmd/main/main.go 

run: 
	go run cmd/main/main.go

watch-css: 
	npx tailwindcss -o static/styles.css --watch
