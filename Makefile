BINARY_NAME=portfolio

all:
	make build-css && make build

build: 
	go build -o ./${BINARY_NAME} ./cmd/portfolio/main.go 

clean:
	rm ./${BINARY_NAME}

run: 
	go run cmd/portfolio/main.go

watch-css: 
	npx tailwindcss -i styles/styles.css -o static/styles/styles.css --watch

build-css: 
	npx tailwindcss -i styles/styles.css -o static/styles/styles.css
