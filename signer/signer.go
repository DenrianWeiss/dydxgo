package signer

type CryptoSigner func(digest []byte) ([]byte, error)
