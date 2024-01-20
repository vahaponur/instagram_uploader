package instagram_uploader

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// UploadSinglePost Creates a container after then, publishes it,if successful returns the IG MediaID
func (receiver IGUploader) UploadSinglePost(post SinglePost) (string, error) {
	endpoint := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/media", receiver.ID)
	containerParams := url.Values{}
	containerParams.Add("access_token", receiver.AccessToken)
	containerParams.Add("caption", post.Caption)
	containerParams.Add("image_url", post.ImageURL)
	containerRes, err := http.PostForm(endpoint, containerParams)
	if err != nil {
		return "", err
	}
	err = containerRes.Body.Close()
	if err != nil {
		return "", err
	}

	var containerResponse IGContainerResponse
	err = json.NewDecoder(containerRes.Body).Decode(&containerResponse)
	if err != nil {
		return "", err
	}
	publishParams := url.Values{}
	publishParams.Add("creation_id", containerResponse.ContainerID)
	publishParams.Add("access_token", receiver.AccessToken)
	publishRes, err := http.PostForm(endpoint, publishParams)
	if err != nil {
		return "", err
	}
	var publishResponse IGPublishResponse
	err = json.NewDecoder(publishRes.Body).Decode(&publishResponse)
	if err != nil {
		return "", err
	}
	return publishResponse.MediaID, nil

}