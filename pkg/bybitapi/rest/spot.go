package rest

import (
	"net/http"
)

// GetWalletBalanceSpot
func (b *ByBit) GetWalletBalanceSpot() (*string, []byte, *ResultSpot, error) {
	var ret ResultSpot

	query, resp, err := b.SignedRequest(
		http.MethodGet, "/spot/v1/account", make(map[string]interface{}), &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return &query, resp, &ret, nil
}
