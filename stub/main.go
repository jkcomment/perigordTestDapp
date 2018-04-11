// Invokes the perigord driver application

package main

import (
	_ "perigordTestDapp/migrations"
	"github.com/polyswarm/perigord/stub"
)

func main() {
	stub.StubMain()
}
