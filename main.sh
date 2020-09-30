nim c src/commands/src/help.nim && mv src/commands/src/help src/commands/bin
nim c src/commands/src/ls.nim && mv src/commands/src/ls src/commands/bin
clear
go run src/gosh.go