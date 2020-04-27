package ufo

import (
	"github.com/blocktree/beam-adapter/beam"
	"github.com/blocktree/openwallet/log"
)

const (
	Symbol = "UFO"
)

type WalletManager struct {
	*beam.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = beam.NewWalletManager()
	wm.Config = beam.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}
