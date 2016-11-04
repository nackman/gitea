package conf

//go:generate go-bindata -ignore "\\.go" -pkg conf -o bindata.go ./...
//go:generate go fmt bindata.go
