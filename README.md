# SnapCooter - the only agile cooter in the assembly

Turtle graphics is an ubiquitous phenomenon of the computer world, a part of the Logo mission of introducing computer programming to the broadest audience possible. Starting with elementary education, it was meant to branch out and change the perception of programming languages within our culture.

This package uses some of the turtle graphic principles in order to generate unique visualization/chart which can then be used to generate some kinds of data.

## Development

Since the project is in its nativity still, it can only be viewed by setting up the dev environment. This project uses [Go 1.17](https://go.dev/) and is buit into WebAssembly. Frontend is being built with [alpine.js](https://alpinejs.dev/) and is contained in few files in the */web* folder, using cdn for all libraries. Frontend component serves both as a showcase as well as a 'real' frontend, but the idea is that all wasm callbacks are general enough to be used in any frontend.

All dev tasks are automated via [Air](https://github.com/cosmtrek/air) and the .air.toml configuration file. By invoking the command *air* in the project's folder wasm binary will be rebuilt, and all development files will be watched for changes. The app is served at localhost:8080 by default. Also *wasm_exec.js* should be copied over from go env in the */web* folder (*cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./web/*).