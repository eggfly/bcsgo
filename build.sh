go get
go build
go test -v -short -cover -covermode=count -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o cover.html

