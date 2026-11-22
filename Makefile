.PHONY: default build test help

# Default target to build the project
default: build

help:	## List all available commands
	@echo 'Available commands:'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

test:	## Run all tests
	go test ./... -count=1

build:	## Build the project
	go build -x

.PHONY: exe
exe:		## Create complete build of this project
	go build -x	
	@CURRENT_TAG=$$(git tag --points-at | head -n 1);\
	CURRENT_TAG=$$(echo "$${CURRENT_TAG////-}");\
	if [ ! -z "$${CURRENT_TAG}" ]; then\
		mv gf2-dc-bot.exe "gf2-dc-bot-$${CURRENT_TAG}.exe";\
		printf '@echo off\r\nstart "" "gf2-dc-bot-$${CURRENT_TAG}.exe"\r\n' > "start.bat"; \
		echo "Window EXE Distribution ready at 'gf2-dc-bot-$${CURRENT_TAG}.exe'.";\
	else\
		mv gf2-dc-bot.exe "gf2-dc-bot-testing.exe";\
		printf '@echo off\r\nstart "" "gf2-dc-bot-testing.exe"\r\n' > "start.bat"; \
		echo "Window EXE Distribution ready at 'gf2-dc-bot-testing.exe'.";\
	fi