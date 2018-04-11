package main

import (
	"context"
	"encoding/json"
	"log"
	"perigordTestDapp/bindings"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/polyswarm/perigord"
	"github.com/polyswarm/perigord/contract"
)

type Greeter struct {
	session *bindings.GreeterSession
	client  *ethclient.Client
}

func NewGreeter(session *bindings.GreeterSession, client *ethclient.Client) *Greeter {
	return &Greeter{
		session: session,
		client:  client,
	}
}

func (g *Greeter) Greet() string {
	greet, _ := g.session.Greet()
	return greet
}

func (g *Greeter) SetGreeting(greeting string) error {
	if result, _ := g.session.SetGreeting(greeting); result == nil {
		log.Fatalln("failed set greeting.")
	}
	return nil
}

// イベント監視
func (g *Greeter) WatchForEvents(eventChan chan *Event) error {
	topics := map[string]common.Hash{
		"Result": perigord.EventSignatureToTopicHash("Result(address,string)"),
	}

	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("Greeter")},
		Topics: [][]common.Hash{{
			topics["Result"],
		}},
	}

	dec := json.NewDecoder(strings.NewReader(bindings.GreeterABI))
	var abi abi.ABI
	if err := dec.Decode(&abi); err != nil {
		return err
	}

	logChan := make(chan types.Log)
	sub, err := greeter.client.SubscribeFilterLogs(context.Background(), q, logChan)
	if err != nil {
		return err
	}
	go func() {
		log.Println("Starting event monitor")
		for {
			select {
			case logMsg := <-logChan:
				if len(logMsg.Topics) != 1 {
					log.Println("incorrect number of topics")
					break
				}

				var nbe NewEventLog
				if err := abi.Unpack(&nbe, "Result", logMsg.Data); err != nil {
					log.Println("error unpacking log: ", err)
					break
				}

				event := &Event{
					Type: "Result",
					Body: NewEventFromLog(nbe),
				}
				eventChan <- event
				break
			case err := <-sub.Err():
				log.Println(err)
				break
			}
		}
	}()

	return nil
}
