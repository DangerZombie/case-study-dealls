package request

type BuySubscriptionRequest struct {
	Body BuySubscriptionRequestBody
}

type BuySubscriptionRequestBody struct {
	Id string
}

type CreateUserInteractionRequest struct {
	UserId1         string `json:"user_id_sender"`
	UserId2         string `json:"user_id_receiver"`
	InteractionType string `json:"interaction_type"`
}

type GetUserToSwipeRequest struct {
	Id string `json:"id"`
}

type LoginRequest struct {
	Body LoginRequestBody `json:"body"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Body RegisterUserRequestBody `json:"body"`
}

type RegisterUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

type ResetSwipeCountRequest struct {
	Status     string
	SwipeCount int
}
