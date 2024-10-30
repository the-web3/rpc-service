package client

type SupportChainResponse struct {
	Support bool `json:"support"`
}

type WalletAddressResponse struct {
	PublicKey string `json:"publicKey"`
	Address   string `json:"address"`
}
