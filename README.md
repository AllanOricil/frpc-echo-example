# Simplest fRPC Example

Contains the source code for the example explained in the [fRPC Quick Start Guide](https://frpc.io/getting-started/quick-start).
This one actually works...


## Requirements

- go 		>= 1.18
- protoc 	>= 3

```bash
brew install go
brew install protoc
```

## Go setup (optional if you have already done it)

Create a directory to serve as the workspace for your go projects

```bash
mkdir $HOME/go
```

Setup `GOPATH`, `GOBIN` and `PATH` env variables as follows

```bash
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

Ensure these variables are also present on every new shell session running the following code to change your shell rc file:

```bash
echo "export GOPATH=$HOME/go" >> $HOME/.zshrc
echo "export GOBIN=$GOPATH/bin" >> $HOME/.zshrc
echo "export PATH=$PATH:$GOBIN" >> $HOME/.zshrc
```

## 1 - Install protoc-gen-go-frpc

Install `protoc-gen-go-frpc` plugin running the following command:

```bash
go install github.com/loopholelabs/frpc-go/protoc-gen-go-frpc@latest
```

Verify the package was installed correctly running the following command:

```bash
which protoc-gen-go-frpc
```

From the root of this project, run the following comand to generate the server and client code using the specs in `echo.proto`

```bash
protoc --go-frpc_out=. echo.proto
```

Verify that `echo/echo.frpc.go` was created in the root of the project. Open it and change `github.com/loopholelabs/polyglot-go` to `github.com/loopholelabs/polyglot` in the imports section, and save.

## 2 - Install project dependencies

```bash
go mod download
```

## 3 - Run Server

Open a new terminal and start the server running the following command

```bash
go run server/main.go
```

![Server](./images/server.png)


## 4 - Run the Client

Open a new terminal and start the client running the following command

```bash
go run client/main.go
```

![Client](./images/client.png)

