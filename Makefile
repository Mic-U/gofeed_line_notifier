clean:
	rm -rf ./dist/* 
build:
	go build -a -tags netgo -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o ./dist/exec ./src