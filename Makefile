build:
	go build -o ./bin/streepjes ./cmd/streepjes/

generate:
	go generate ./...

resetdb:
	rm streepjes.db -f
	cp streepjes.example.db streepjes.db

newtestdb:
	rm streepjes.db -f
	go run ./cmd/maketestdb/
	cp streepjes.db streepjes.example.db

run:
	go run ./cmd/streepjes/