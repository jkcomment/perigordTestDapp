package main

import "github.com/ethereum/go-ethereum/common"

type Event struct {
	Type string
	Body interface{}
}

type NewEventLog struct {
	From   common.Address
	Stored string
}

type NewEvent struct {
	From   common.Address
	Stored string
}

func NewEventFromLog(nbe NewEventLog) *NewEvent {
	return &NewEvent{
		From:   nbe.From,
		Stored: nbe.Stored,
	}
}
