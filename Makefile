test:
	go test ./... --cover -coverprofile=cover.out -v

run:
	go run main.go