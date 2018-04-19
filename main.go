package main

import (
	"context"
	"fmt"
	"log"
	"perigordTestDapp/bindings"
	_ "perigordTestDapp/migrations"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"
)

var greeter *Greeter

func main() {
	// 初期化
	network.InitNetworks()

	// perigord.yamlから情報を取得し、接続を行う
	nw, err := network.Dial("dev")
	if err != nil {
		log.Fatalln("Could not connect to dev network: ", err)
	}

	// マイグレーション実行(デプロイ)
	if err := migration.RunMigrations(context.Background(), nw, true); err != nil {
		log.Fatalln("Error running migrations: ", err)
	}

	// コントラクト接続
	session, ok := contract.Session("Greeter").(*bindings.GreeterSession)
	if !ok {
		log.Fatalln("Error retrieving session")
	}

	greeter = NewGreeter(session, nw.Client())
	// イベント監視用チャンネル
	eventChan := make(chan *Event)
	if err := greeter.WatchForEvents(eventChan); err != nil {
		log.Println("error listening for incoming events:", err)
		return
	}

	// データ取得
	fmt.Printf("更新前: %s\n", greeter.Greet())

	// データ更新
	greeter.SetGreeting("Hello, World!:D")

	//  イベント監視
	for {
		event := <-eventChan
		fmt.Println("Event結果\n")
		fmt.Printf("from: %v\n", event.Body.(*NewEvent).From.Hex())
		fmt.Printf("stored: %v\n", event.Body.(*NewEvent).Stored)
		// 先ほど更新した内容("Hello, World!:D")に更新されているか確認
		fmt.Printf("更新後: %s\n", greeter.Greet())
	}
}
