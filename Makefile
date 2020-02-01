BINARY_NAME=npmdepcopy

editor.exe:	**/*.go
	mkdir -p builds
	go build GOOS=windows GOARCH=amd64 -o builds/$(BINARY_NAME).exe -v ./editor/...