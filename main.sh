nim c src/commands/src/help.nim && mv src/commands/src/help src/commands/bin
nim c src/commands/src/calc.nim && mv src/commands/src/calc src/commands/bin
clear
go run src/gosh.go