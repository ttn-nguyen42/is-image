package responses

/*
 * /api/v1/upload_image
 */
type PostUploadImageResponse struct {
	Result string `json:"result"`
	Type   string `json:"image_type"`
	HashID string `json:"image_id"`
}
