package main

import (
	"nps/bridge"
	"sync"
)

var (
	Bridge  *bridge.Bridge
	RunList sync.Map //map[int]interface{}
)

func init() {
	RunList = sync.Map{}
}