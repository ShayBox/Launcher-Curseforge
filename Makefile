all: clean darwin linux-386 linux-amd64 windows

clean:
	-rm -rf ./output
	-rm ./build/darwin/MultiMC-Curseforge.app/Contents/MacOS/MultiMC-Curseforge
	-rm ./build/linux/usr/bin/multimc-curseforge

darwin:
	mkdir -p output
	GOOS=darwin GOARCH=amd64 go build -o ./build/darwin/MultiMC-Curseforge.app/Contents/MacOS/MultiMC-Curseforge
	tar -czvf ./output/darwin.tar.gz --exclude .gitkeep -C ./build/darwin .

linux-386:
	mkdir -p output
	GOOS=linux GOARCH=386 go build -o ./build/linux/usr/bin/multimc-curseforge
	tar -czvf ./output/linux-386.tar.gz --exclude .gitkeep -C ./build/linux .

linux-amd64:
	mkdir -p output
	GOOS=linux GOARCH=amd64 go build -o ./build/linux/usr/bin/multimc-curseforge
	tar -czvf ./output/linux-amd64.tar.gz --exclude .gitkeep -C ./build/linux .

windows:
	mkdir -p output
	GOOS=windows GOARCH=386 go build -o ./output/MultiMC-Curseforge.exe