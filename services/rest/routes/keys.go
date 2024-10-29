package routes

import (
	"fmt"
	models "github.com/the-web3/rpc-service/services/rest/model"
	"net/http"
)

func (h Routes) GetSupportCoins(w http.ResponseWriter, r *http.Request) {
	chain := r.URL.Query().Get("chain")
	network := r.URL.Query().Get("network")
	cr := &models.ChainRequest{
		Chain:   chain,
		Network: network,
	}
	supRet, err := h.svc.GetSupportCoins(cr)
	if err != nil {
		return
	}
	err = jsonResponse(w, supRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}

func (h Routes) GetWalletAddress(w http.ResponseWriter, r *http.Request) {
	chain := r.URL.Query().Get("chain")
	network := r.URL.Query().Get("network")
	cr := &models.ChainRequest{
		Chain:   chain,
		Network: network,
	}

	addrRet, err := h.svc.GetWalletAddress(cr)
	if err != nil {
		return
	}

	err = jsonResponse(w, addrRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}
