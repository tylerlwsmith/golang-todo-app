# Go Todo App

I am building a small todo app as an opportunity to learn Go. This is a work-in-progress and likely written in unidiomatic Go.

## Install locally

After cloning the repo, run the following command to install the Go dependencies locally:

```sh
go mod tidy
```

This project also uses the `prettier` JavaScript package with the `prettier-plugin-go-template` plugin to format go templates. Run the following command to install these for template formatting:

```sh
npm install
```

Live reloading of the server during development is accomplished with the `air` utility. You can install it with the following command:

```sh
go install github.com/cosmtrek/air@latest
```

After installing `air`, your machine may not have the binary in the `PATH`. You can create an alias in `.bashrc` or appropriate file where the system's shell settings are loaded:

```txt
alias air='$(go env GOPATH)/bin/air'
```

If you're developing with VS Code and you want better support for `.tmpl` files, add the following to your `settings.json` file:

```json
"files.associations": {
  "*.tmpl": "html"
}
```

## Issues

The Go language server has trouble with build tags when developing on Linux. You can see the [LSP GitHub issue here](https://github.com/golang/go/issues/29202).
