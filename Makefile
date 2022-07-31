test:
	go test ./... --cover -coverprofile=coverage.out -v

run:
	go run main.go