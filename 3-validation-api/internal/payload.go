package internal

type SendPayload struct {
	Email string `json:"email"`
}

type VerifyPayload struct {
	Hash string `json:"hash"`
}
