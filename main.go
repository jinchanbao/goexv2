package goex

import (
	"reflect"

	"github.com/onebitorg/goex/binance"
	"github.com/onebitorg/goex/httpcli"
	"github.com/onebitorg/goex/huobi"
	"github.com/onebitorg/goex/logger"
	"github.com/onebitorg/goex/okx"
)

var (
	DefaultHttpCli = httpcli.Cli
)

var (
	OKx     = okx.New()
	Binance = binance.New()
	HuoBi   = huobi.New()
)

func SetDefaultHttpCli(cli httpcli.IHttpClient) {
	logger.Infof("use new http client implement: %s", reflect.TypeOf(cli).Elem().String())
	httpcli.Cli = cli
}
