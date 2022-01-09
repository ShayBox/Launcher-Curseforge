GOCMD=go
GOFLAGS=-o trace -gcflags "all=-trimpath=$GOPATH"

all: clean darwin linux windows

clean:
	-rm -rf ./output
	-rm ./build/darwin/PolyMC-Curseforge.app/Contents/MacOS/PolyMC-Curseforge
	-rm ./build/linux/usr/bin/polymc-curseforge

darwin:
	mkdir -p output
	GOOS=darwin GOARCH=amd64 $(GOCMD) build $(GOFLAGS) -o ./build/darwin/PolyMC-Curseforge.app/Contents/MacOS/PolyMC-Curseforge
	tar -czvf ./output/darwin.tar.gz --exclude .gitkeep -C ./build/darwin .

linux:
	mkdir -p output
	GOOS=linux GOARCH=386 $(GOCMD) build $(GOFLAGS) -o ./build/linux/usr/bin/polymc-curseforge
	tar -czvf ./output/linux.tar.gz --exclude .gitkeep -C ./build/linux .

windows:
	mkdir -p output
	GOOS=windows GOARCH=386 $(GOCMD) build $(GOFLAGS) -o ./output/PolyMC-Curseforge.exe