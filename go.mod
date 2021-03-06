module github.com/assetsadapterstore/ufo-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.11.1
	github.com/blocktree/beam-adapter v1.1.1
	github.com/blocktree/go-owcrypt v1.0.1
	github.com/blocktree/openwallet v1.5.2
	github.com/imroc/req v0.2.3
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/tidwall/gjson v1.2.1
	gopkg.in/urfave/cli.v1 v1.20.0
)

//replace github.com/blocktree/beam-adapter => ../beam-adapter
