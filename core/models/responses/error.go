package responses

/*
 * Response message sent back on errors
 */
type ErrorResponse struct {
	Message string `json:"message"`
}
