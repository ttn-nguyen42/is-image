package model

/*
 * Response message sent back on errors
 */
type ErrorResponse struct {
	Message string `json:"message"`
}
