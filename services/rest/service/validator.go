package service

type Validator struct{}

func (v *Validator) VerifyWalletAddress(chain, network string) bool {
	if chain != "Bitcoin" && chain != "Ethereum" {
		return false
	}
	if network != "MainNet" && network != "TestNet" {
		return false
	}
	return true
}
