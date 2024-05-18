package response

type BuySubscriptionResponse struct {
	Message string `json:"message"`
}

type CreateUserInteractionResponse struct {
	Message string `json:"message"`
}

type GetUserToSwipeResponse struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type ResetSwipeCountResponse struct {
	Message string `json:"message"`
}
