package instagram_uploader

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// UploadCarouselImagePost Uploads Carousel (multiple) Image post, if no error, returns media id
func (receiver IGUploader) UploadCarouselImagePost(post CarouselImagePost) (string, error) {
	containerEndpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/media", receiver.Version, receiver.ID)
	containerParams := url.Values{}
	containerParams.Add("is_carousel_item", "true")
	containerParams.Add("access_token", receiver.AccessToken)
	imageContainers := make([]string, 0)

	for _, imageURL := range post.ImageURLS {
		containerParams.Add("image_url", imageURL)
		containerRes, err := http.PostForm(containerEndpoint, containerParams)
		if err != nil {
			return " ", err
		}
		var imageContainerResponse IGContainerResponse
		err = json.NewDecoder(containerRes.Body).Decode(&imageContainerResponse)
		if err != nil {
			return "", err
		}

		imageContainers = append(imageContainers, imageContainerResponse.ContainerID)
		containerParams.Del("image_url")
		err = containerRes.Body.Close()
		if err != nil {
			return "", err
		}
	}

	children := strings.Join(imageContainers, ",")

	carouselParams := url.Values{}
	carouselParams.Add("media_type", "CAROUSEL")
	carouselParams.Add("children", children)
	carouselParams.Add("caption", post.Caption)
	carouselParams.Add("access_token", receiver.AccessToken)

	carouselContainerRes, err := http.PostForm(containerEndpoint, carouselParams)
	if err != nil {
		return "", err
	}
	defer carouselContainerRes.Body.Close()
	var carouselResponse IGContainerResponse
	err = json.NewDecoder(carouselContainerRes.Body).Decode(&carouselResponse)
	if err != nil {
		return "", err
	}

	publishEndpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/media_publish", receiver.Version, receiver.ID)
	publishParams := url.Values{}
	publishParams.Add("creation_id", carouselResponse.ContainerID)
	publishParams.Add("access_token", receiver.AccessToken)
	publishRes, err := http.PostForm(publishEndpoint, publishParams)
	defer publishRes.Body.Close()
	if err != nil {
		return "", err
	}
	var publishResponse IGPublishResponse
	err = json.NewDecoder(publishRes.Body).Decode(&publishResponse)
	if err != nil {
		return "", err
	}
	mediaID := publishResponse.MediaID
	return mediaID, nil
}
