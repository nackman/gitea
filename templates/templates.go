package templates

//go:generate go-bindata -ignore "\\.go" -pkg templates -o bindata.go ./...
//go:generate go fmt bindata.go
