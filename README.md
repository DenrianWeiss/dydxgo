# DyDxGo

DyDx client for Golang, some of the code (most of private/public part) taken from [go-dydx](https://github.com/go-numb/go-dydx)  
With Support for custom local signer and on-chain txs

## Notice

### OnBoarding Namespace

DyDx's onboarding system strongly depends on RFC6979, which IS NOT implemented on some signers

Supported:

- Ledger
- Trezor
- Most Browser Wallets
- GoEthereum's crypto.sign()

NOT Supported:

- AWS KMS
- GCP's cloud key management

### Ethereum Namespace

We use GoEthereum's ethclient interface to do this part, those key management service not supported by DyDx's key derivation system still can be used here.

