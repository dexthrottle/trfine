package rest

import "net/http"

// GetWalletBalanceSpot
func (b *ByBit) GetWalletBalanceSpot() (query string, resp []byte, result ResultSpot, err error) {
	var ret ResultSpot

	query, resp, err = b.SignedRequest(http.MethodGet, "/spot/v1/account", nil, &ret)
	if err != nil {
		return
	}

	return
}
