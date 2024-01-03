package tests

import (
	"crypto/ecdsa"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/denrianweiss/dydxgo/client"
	"github.com/denrianweiss/dydxgo/client/private"
	"github.com/denrianweiss/dydxgo/client/public"
	"github.com/denrianweiss/dydxgo/signer"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func GetSigner(privateKey string) signer.CryptoSigner {
	curve := secp256k1.S256()
	pkNumber, _ := big.NewInt(1).SetString(strings.TrimPrefix(privateKey, "0x"), 16)
	x, y := curve.ScalarBaseMult(pkNumber.Bytes())

	pk := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: pkNumber,
	}

	return func(digest []byte) ([]byte, error) {
		return crypto.Sign(digest, &pk)
	}
}

func RandomClientId() string {
	// Generate Random String
	random, _ := uuid.NewRandom()
	return random.String()
}

func GetDateString(validDuration time.Duration) string {
	return time.Now().Add(validDuration).UTC().Format("2006-01-02T15:04:05.999Z")
}

func TestPrivateApi(t *testing.T) {
	// Read Private Key from env.
	pk := os.Getenv("PRIVATE_KEY")
	if pk == "" {
		t.Fatal("Please set PRIVATE_KEY env.")
	}
	address := os.Getenv("ADDRESS")
	// Get Signer.
	cryptoSigner := GetSigner(pk)
	// Get Client.
	logger := log.New(os.Stdout, "dydxgo: ", log.LstdFlags)
	// Do OnBoarding.
	dydxClient := client.New(client.Options{
		NetworkId:              11155111,
		Host:                   "https://api.stage.dydx.exchange",
		DefaultEthereumAddress: address,
		StarkPrivateKey:        "",
		ApiKeyCredentials:      nil,
		HttpClient:             http.DefaultClient,
		Logger:                 logger,
		CryptoSigner:           cryptoSigner,
		EthSigner:              nil,
		EthConn:                nil,
	})
	// Get Account.
	acc, err := dydxClient.Private.GetAccount(address)
	if err != nil {
		t.Fatal(err.Error())
	}
	// Get Market
	orders, err := dydxClient.Public.GetOrderBook("ETH-USD")
	if err != nil {
		t.Fatal(err.Error())
	}
	reg, err := dydxClient.Private.GetRegistration()
	if err != nil {
		t.Fatal(err.Error())
	}
	// Print Registration
	log.Printf("Registration: %v", reg)
	// Create FOK Buy order
	po := &private.ApiOrder{
		Market:      "ETH-USD",
		Side:        "BUY",
		Type:        "LIMIT",
		Size:        "0.01",
		Price:       orders.Bids[1].Price,
		TimeInForce: "FOK",
		LimitFee:    "0.0005",
		PostOnly:    false,
		ClientId:    RandomClientId(),
	}
	po.Expiration = GetDateString(time.Minute * 8)
	order, err := dydxClient.Private.CreateOrder(po, acc.Account.PositionId)
	if err != nil {
		t.Fatal(err.Error())
	}
	// Try Get Order Status
	orderStatus, err := dydxClient.Private.GetOrderById(order.Order.ID)
	if err != nil {
		t.Fatal(err.Error())
	}
	// Check order status
	if orderStatus.Order.Status != "FILLED" {
		log.Print("Order not filled")
	} else {
		// Create sell order
		po := &private.ApiOrder{
			Market:      "ETH-USD",
			Side:        "SELL",
			Type:        "LIMIT",
			Size:        "0.01",
			Price:       orders.Asks[1].Price,
			TimeInForce: "FOK",
			LimitFee:    "0.0005",
			PostOnly:    false,
			ClientId:    RandomClientId(),
		}
		po.Expiration = GetDateString(time.Minute * 8)
		_, err = dydxClient.Private.CreateOrder(po, acc.Account.PositionId)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	// Test Fast Withdraw
	wd, err := dydxClient.Public.GetFastWithdrawals(&public.FastWithdrawParam{CreditAsset: "USDC", CreditAmount: "100"})
	if err != nil {
		t.Fatal(err.Error())
	}
	// Initiate Fast Withdraw
	/// Get First Liquidity provider
	liqId := ""
	for id, _ := range wd.LiquidityProviders {
		liqId = id
		return
	}
	lp := wd.LiquidityProviders[liqId]

	wf, err := dydxClient.Private.WithdrawFast(
		&private.WithdrawalFastParam{
			ClientID:     RandomClientId(),
			ToAddress:    address,
			CreditAsset:  "USDC",
			CreditAmount: "100",
			DebitAmount:  lp.Quote.DebitAmount,
			LpPositionId: liqId,
			Expiration:   GetDateString(8 * 24 * time.Hour),
			Signature:    "",
		}, lp.StarkKey, strconv.FormatInt(acc.Account.PositionId, 10))
	// Check withdraw requests
	if err != nil {
		t.Fatal(err.Error())
	}
	// Get All Transfer
	transfers, err := dydxClient.Private.GetTransfers(&private.TransferQueryParam{})
	if err != nil {
		t.Fatal(err.Error())
	}
	// Find Transfer in it
	found := false
	for _, t := range transfers.Transfer {
		if t.Id == wf.Withdrawal.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Withdraw not found")
	}
	// Test Slow Withdraw
	sw, err := dydxClient.Private.Withdraw(&private.WithdrawalParam{
		ClientID:   RandomClientId(),
		Amount:     "100",
		Asset:      "USDC",
		PositionId: strconv.FormatInt(acc.Account.PositionId, 10),
		Expiration: GetDateString(8 * 24 * time.Hour),
		Signature:  "",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	// Get All Transfer
	transfers, err = dydxClient.Private.GetTransfers(&private.TransferQueryParam{})
	if err != nil {
		t.Fatal(err.Error())
	}
	// Find Transfer in it
	found = false
	for _, t := range transfers.Transfer {
		if t.Id == sw.Withdrawal.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Withdraw not found")
	}
}
