package tests

import (
	"github.com/denrianweiss/dydxgo/client/public"
	"github.com/denrianweiss/dydxgo/types"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestGetOrders(t *testing.T) {
	publicClient := public.Public{}
	logP := log.New(os.Stdout, "dydxgo: ", log.LstdFlags)
	publicClient.Logger = logP
	publicClient.HttpClient = http.DefaultClient
	publicClient.Host = "https://api.stage.dydx.exchange"
	publicClient.NetworkId = 5
	publicClient.RateLimit = &types.RateLimit{}
	// Test Get Markets
	markets, err := publicClient.GetMarkets("ETH-USD")
	if err != nil {
		t.Fail()
	}
	log.Println(markets)
	// Test Get Withdraw.
	withdraw, err := publicClient.GetFastWithdrawals(&public.FastWithdrawParam{CreditAsset: "USDC", CreditAmount: "100"})
	if err != nil {
		t.Fail()
	}
	log.Println(withdraw)
	// Test Get Orders.
	orders, err := publicClient.GetOrderBook("ETH-USD")
	if err != nil {
		t.Fail()
	}
	log.Println(orders)
}
