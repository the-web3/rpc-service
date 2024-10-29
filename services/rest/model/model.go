package models

type ChainRequest struct {
	Chain   string `json:"chain"`
	Network string `json:"network"`
}

type SupportChainResponse struct {
	Support bool `json:"support"`
}

type WalletAddressResponse struct {
	PublicKey string `json:"publicKey"`
	Address   string `json:"address"`
}
