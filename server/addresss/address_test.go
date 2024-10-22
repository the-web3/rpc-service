package addresss

import (
	"fmt"
	"testing"
)

func TestCreateAddressByKeyPairs(t *testing.T) {
	address, err := CreateAddressFromPrivateKey()
	if err != nil {
		return
	}
	fmt.Println(address.PrivateKey, address.PublicKey, address.Address)
}

//	 {
//			"privateKey":"0xfe13c8e55444107c32f50cca04965f9772fe4fd720ffedd30b347e541fe7a97c",
//			"publicKey":"0x03fa3af3ad7a5c97e3b6bd8bc6dd5751c4b8ced139a2b3a62716f70560cf4a211e",
//			"address":"0x35096AD62E57e86032a3Bb35aDaCF2240d55421D"
//	 }
func TestPublicKeyToAddress(t *testing.T) {
	address, err := PublicKeyToAddress("03fa3af3ad7a5c97e3b6bd8bc6dd5751c4b8ced139a2b3a62716f70560cf4a211e")
	if err != nil {
		return
	}
	fmt.Println("address==", address)
}
