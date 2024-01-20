package instagram_uploader

type IGUploader struct {
	ID          string `json:"id"`
	AccessToken string `json:"accessToken"`
}
type SinglePost struct {
	ImageURL string
	Caption  string
}
type IGContainerResponse struct {
	ContainerID string `json:"id"`
}
type IGPublishResponse struct {
	MediaID string `json:"id"`
}
