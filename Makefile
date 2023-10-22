app:
	@go build -o ./cmd/app/bin ./cmd/app/main.go && ./cmd/app/bin/main

api:
	@go build -o ./cmd/api/bin ./cmd/api/main.go && ./cmd/api/bin/main
