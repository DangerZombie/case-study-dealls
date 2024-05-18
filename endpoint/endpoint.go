package endpoint

import "github.com/DangerZombie/case-study-dealls/helper/auth"

type endpointImpl struct {
	authHelper auth.AuthHelper
}

func NewEndpoint(ah auth.AuthHelper) Endpoint {
	return &endpointImpl{ah}
}
