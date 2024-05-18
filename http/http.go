package http

import "github.com/DangerZombie/case-study-dealls/helper/auth"

type httpImpl struct {
	authHelper auth.AuthHelper
}

func NewHttp(ah auth.AuthHelper) Http {
	return &httpImpl{ah}
}
