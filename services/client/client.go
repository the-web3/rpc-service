package client

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
)

var errWaletHTTPError = errors.New("wallet http error")

type Address struct {
	PublicKey string
	Address   string
}

type WalletClient interface {
	GetSupportCoins(chain, network string) (bool, error)
	GetWalletAddress(chain, network string) (*Address, error)
}

type Client struct {
	client *resty.Client
}

func NewWalletClient(url string) *Client {
	client := resty.New()
	client.SetBaseURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			baseUrl := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, baseUrl, errWaletHTTPError)
		}
		fmt.Println("baseUrl::", r.Request.Method)
		fmt.Println("method::", r.Request.URL)
		fmt.Println("method::", r.Request.QueryParam)
		return nil
	})
	return &Client{
		client: client,
	}
}

func (c *Client) GetSupportCoins(chain, network string) (bool, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"chain":   chain,
		"network": network,
	}).SetResult(&SupportChainResponse{}).Get("/api/v1/support_chain")
	if err != nil {
		return false, errors.New("support chain request fail")
	}
	spt, ok := res.Result().(*SupportChainResponse)
	if !ok {
		return false, errors.New("support chain transfer type fail")
	}
	return spt.Support, nil
}

func (c *Client) GetWalletAddress(chain, network string) (*Address, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"chain":   chain,
		"network": network,
	}).SetResult(&WalletAddressResponse{}).Get("/api/v1/wallet_address")
	if err != nil {
		return nil, errors.New("wallet address request fail")
	}
	wap, ok := res.Result().(*WalletAddressResponse)
	if !ok {
		return nil, errors.New("wallet address transfer type fail")
	}
	return &Address{
		PublicKey: wap.PublicKey,
		Address:   wap.Address,
	}, nil
}
