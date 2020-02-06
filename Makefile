all: clean darwin linux windows

clean:
	-rm -rf ./output
	-rm ./build/darwin/MultiMC-Twitch.app/Contents/MacOS/MultiMC-Twitch
	-rm ./build/linux/usr/bin/multimc-twitch
	-rm ./build/windows/MultiMC-Twitch.exe

darwin:
	mkdir -p output
	GOOS=darwin go build -o ./build/darwin/MultiMC-Twitch.app/Contents/MacOS/MultiMC-Twitch
	tar -czvf ./output/darwin.tar.gz --exclude .gitkeep -C ./build/darwin .

linux:
	mkdir -p output
	GOOS=linux go build -o ./build/linux/usr/bin/multimc-twitch
	tar -czvf ./output/linux.tar.gz --exclude .gitkeep -C ./build/linux .

windows:
	mkdir -p output
	GOOS=windows go build -o ./output/MultiMC-Twitch.exe