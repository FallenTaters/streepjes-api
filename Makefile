build:
	go build -o ./bin/streepjes ./cmd/streepjes/

generate:
	go generate ./...
	go-bindata -prefix shared/migrate/files/ -pkg migrate -o ./shared/migrate/bindata.go ./shared/migrate/files/
	go-bindata -prefix cmd/maketestdb/files/ -o ./cmd/maketestdb/bindata.go ./cmd/maketestdb/files/

resetdb:
	rm streepjes.db -f
	cp streepjes.example.db streepjes.db

newtestdb:
	rm streepjes.db -f
	go run ./cmd/maketestdb/
	go run ./cmd/makeadmin/
	cp streepjes.db streepjes.example.db

admin:
	go run ./cmd/makeadmin/

testdata:
	go run ./cmd/maketestdb/