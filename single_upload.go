package instagram_uploader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// UploadSingleImagePost Creates a single(image) media container after then, publishes it,if successful returns the IG MediaID
func (receiver IGUploader) UploadSingleImagePost(post SingleImagePost) (string, error) {
	containerEndpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/media", receiver.Version, receiver.ID)
	containerParams := url.Values{}
	containerParams.Add("access_token", receiver.AccessToken)
	containerParams.Add("caption", post.Caption)
	containerParams.Add("image_url", post.ImageURL)
	containerRes, err := http.PostForm(containerEndpoint, containerParams)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(containerRes.Body)

	if err != nil {
		return "", err
	}

	var containerResponse IGContainerResponse
	err = json.NewDecoder(containerRes.Body).Decode(&containerResponse)
	if err != nil {
		return "", err
	}
	publishEndpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/media_publish", receiver.Version, receiver.ID)
	publishParams := url.Values{}
	publishParams.Add("creation_id", containerResponse.ContainerID)
	publishParams.Add("access_token", receiver.AccessToken)
	publishRes, err := http.PostForm(publishEndpoint, publishParams)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println()
		}
	}(publishRes.Body)
	if err != nil {
		return "", err
	}
	var publishResponse IGPublishResponse
	err = json.NewDecoder(publishRes.Body).Decode(&publishResponse)
	if err != nil {
		return "", err
	}
	mediaId := publishResponse.MediaID
	return mediaId, nil

}
