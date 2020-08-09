package Models

/**
 * LoginResponse
 * Is the model to response with jwt
 */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}