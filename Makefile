all: tailwind server

tests:
	go test -v ./...

tailwind:
	@cd views; npx tailwindcss -i ./../static/css/index.css -o ./../static/css/styles.css --watch

server:
	@air
