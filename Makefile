build:
	@go build -o bin

run: build
	@./bin/https-server
	explain me this code:
This Makefile contains two targets: `build` and `run`.
- The `build` target compiles the Go source code into an executable binary named `https-server` and places it in the `bin` directory. The `@` symbol before the command suppresses the command output in the terminal.
what does @ do in makefile?
- The `run` target first depends on the `build` target, meaning it will execute the `build` target before running its own commands. After building the binary, it runs the `https-server` executable located in the `bin` directory. Again, the `@` symbol is used to suppress the command output.
explain in simple what is @ for?
The `@` symbol in a Makefile is used to prevent the command from being printed to the terminal when it is executed. This makes the output cleaner by only showing the results of the command, not the command itself.