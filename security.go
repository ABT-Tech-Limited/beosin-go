package beosin

import "context"

const (
	endpointBlackScreening = "/api/v1/tag/black/screening"
)

// BlackAddressScreening performs black address screening
func (c *client) BlackAddressScreening(ctx context.Context, req *BlackScreeningRequest) (*BlackScreeningResponse, error) {
	params := buildQueryParams(map[string]string{
		"platform": req.Platform,
		"address":  req.Address,
	})

	var resp BlackScreeningResponse
	if err := c.doRequest(ctx, endpointBlackScreening, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
