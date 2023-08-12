package main

import (
	"GopherSentinel/BotOperations/WatchChannel"
)

func getBotDependencies() *WatchChannel.WatchChannel {
	return &WatchChannel.WatchChannel{}
}

func main() {
	op := getBotDependencies()
	op.StartBot()
}
