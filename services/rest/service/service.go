package service

import (
	"github.com/the-web3/rpc-service/database"
	models "github.com/the-web3/rpc-service/services/rest/model"
)

type Service interface {
	GetSupportCoins(*models.ChainRequest) (*models.SupportChainResponse, error)
	GetWalletAddress(*models.ChainRequest) (*models.WalletAddressResponse, error)
}

type HandleSrv struct {
	v        *Validator
	keysView database.KeysView
}

func NewHandleSrv(v *Validator, ksv database.KeysView) Service {
	return &HandleSrv{
		v:        v,
		keysView: ksv,
	}
}

func (h HandleSrv) GetSupportCoins(req *models.ChainRequest) (*models.SupportChainResponse, error) {
	ok := h.v.VerifyWalletAddress(req.Chain, req.Network)
	if ok {
		return &models.SupportChainResponse{
			Support: true,
		}, nil
	} else {
		return &models.SupportChainResponse{
			Support: false,
		}, nil
	}
}

func (h HandleSrv) GetWalletAddress(*models.ChainRequest) (*models.WalletAddressResponse, error) {
	return &models.WalletAddressResponse{
		PublicKey: "public key",
		Address:   "0x00000000000000000000000",
	}, nil
}
