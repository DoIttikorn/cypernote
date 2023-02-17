install:
	go mod tidy

dev:
	DB_CONNECTION=postgresql://postgres@localhost:5432/cybernote?sslmode=disable go run .