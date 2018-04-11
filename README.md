# This is Perigord's tutorial.

## What is Perigord?
Perigord is Golang Tools for Ethereum Development.

Github: [polyswarm/perigord](https://github.com/polyswarm/perigord)

Please refer here for the installation. Click [here](https://github.com/polyswarm/perigord)

## Install
```
$ cd $GOPATH/src && git clone https://github.com/jkcomment/perigordTestDapp
```

## Build & Run
Before build there is a testnet's script in the scripts directory in the Perigord repository. Execute testnet's script on another terminal.

```
$ $GOPATH/src/github.com/polyswarm/perigord/scripts/launch_geth_testnet.sh
```

build & run
```
$ cd perigordTestDapp && go run main.go event.go greeter.go
$ ./main
```
