package instagram_uploader

type IGUploader struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Version     string `json:"version"`
}
type SingleImagePost struct {
	ImageURL string `json:"image_url"`
	Caption  string `json:"caption"`
}
type IGContainerResponse struct {
	ContainerID string `json:"id"`
}
type IGPublishResponse struct {
	MediaID string `json:"id"`
}
type CarouselImagePost struct {
	ImageURLS []string `json:"image_urls"`
	Caption   string   `json:"caption"`
}
