package client

import (
	"fmt"
	"testing"
)

func TestSupportChain(t *testing.T) {
	client := NewWalletClient("http://127.0.0.1:8970")
	result, err := client.GetSupportCoins("Bitcoin", "MainNet")
	if err != nil {
		fmt.Println("get support chain fail")
		return
	}
	fmt.Println("Support Chain Res:", result)
}

func TestWalletAddress(t *testing.T) {
	client := NewWalletClient("http://127.0.0.1:8970")
	addressInfo, err := client.GetWalletAddress("Bitcoin", "MainNet")
	if err != nil {
		fmt.Println("get wallet address fail")
		return
	}
	fmt.Println("Wallet Address Result:", addressInfo.Address, addressInfo.PublicKey)
}

// http://127.0.0.1/wallet/1
// http://127.0.0.1/api/v1?address=111
// http://127.0.0.1/api/v1
/*
{
	"address": "0x00000000"
}
*/
