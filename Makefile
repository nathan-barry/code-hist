run:
	clear && ~/tailwindcss -i views/input.css -o static/tailwind.css && go build -o ./cmd/app/bin ./cmd/app/main.go && ./cmd/app/bin/main

build:
	clear && ~/tailwindcss -i views/input.css -o static/tailwind.css --minify && go build -o ./cmd/app/bin ./cmd/app/main.go 
