package openwtester

import (
	"github.com/assetsadapterstore/ufo-adapter/ufo"
	"github.com/astaxie/beego/config"
	"path/filepath"
	"time"
)

var (
	serverNode *ufo.WalletManager
	clientNode *ufo.WalletManager
)

func init() {

	//serverNode = testNewWalletManager("server.ini")
	clientNode = testNewWalletManager("UFO.ini")
	time.Sleep(1 * time.Second)
}

func testNewWalletManager(conf string) *ufo.WalletManager {
	wm := ufo.NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", conf)
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	return wm
}

