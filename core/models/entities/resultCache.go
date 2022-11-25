package entities

/*
 * Data model of result in cache database
 */
type ResultCache struct {
	Id        string `json:"image_id"`
	Imagetype string `json:"image_type"`
	Result    string `json:"result"`
}
