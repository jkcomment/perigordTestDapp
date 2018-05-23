# This is Perigord's sample code.

## What is Perigord?
Perigord is Golang Tools for Ethereum Development.

Github: [polyswarm/perigord](https://github.com/polyswarm/perigord)

Please refer here for the installation. Click [here](https://github.com/polyswarm/perigord)

## Installation
```
$ cd $GOPATH/src && git clone https://github.com/jkcomment/perigordTestDapp
```

## Usage
First, it is necessary to run testnet before the build. A testnet script that can easily run testnet is in the Perigord repository. Execute the script file.
A testnet script that can easily run testnet is in the Perigord repository. Run it before the build.

```
$ $GOPATH/src/github.com/polyswarm/perigord/scripts/launch_geth_testnet.sh
```

And, build & run
```
$ cd perigordTestDapp && go build main.go event.go greeter.go
$ ./main
```
