run:
	clear && go build -o ./cmd/app/bin ./cmd/app/main.go && ./cmd/app/bin/main

	
tailwind:
	~/tailwindcss -i views/input.css -o static/tailwind.css --watch

build:
	@clear && go build -o ./cmd/app/bin ./cmd/app/main.go && ~/tailwindcss -i views/input.css -o static/tailwind.css --minify
