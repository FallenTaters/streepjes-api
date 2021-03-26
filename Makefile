generate:
	go generate ./...
	go-bindata -pkg migrate -o ./shared/migrate/bindata.go ./shared/migrate/files/
	go-bindata -prefix cmd/maketestdb/files/ -o ./cmd/maketestdb/bindata.go ./cmd/maketestdb/files/

resetdb:
	rm streepjes.db

admin:
	go run ./cmd/makeadmin/

testdata:
	go run ./cmd/maketestdb/