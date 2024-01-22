package instagram_uploader

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

const (
	BaseMediaEndpoint    = "https://graph.facebook.com/v18.0/%s/media"
	MediaPublishEndpoint = "https://graph.facebook.com/v18.0/%s/media_publish"
)

// IGUploader is a struct to handle Instagram uploading functionality.
type IGUploader struct {
	ID          string
	AccessToken string
}

// IGContainerResponse holds the response for media container creation.
type IGContainerResponse struct {
	ContainerID string `json:"id"`
}

// IGPublishResponse holds the response for media publishing.
type IGPublishResponse struct {
	MediaID string `json:"id"`
}

// New creates a new instance of IGUploader.
func New(id string, token string) IGUploader {
	return IGUploader{ID: id, AccessToken: token}
}

// UploadImagePost handles the uploading of single or multiple images to Instagram.
func (u IGUploader) UploadImagePost(images []string, caption string) (string, error) {
	countOfImages := len(images)
	if countOfImages == 0 {
		return "", errors.New("images empty")
	}

	if countOfImages == 1 {
		return u.uploadSingleImagePost(images[0], caption)
	} else {
		return u.uploadCarouselImagePost(images, caption)
	}
}

// uploadSingleImagePost handles uploading a single image to Instagram.
func (u IGUploader) uploadSingleImagePost(image string, caption string) (string, error) {

	client := resty.New()

	containerEndpoint := fmt.Sprintf(BaseMediaEndpoint, u.ID)
	publishEndpoint := fmt.Sprintf(MediaPublishEndpoint, u.ID)

	var containerResponse IGContainerResponse
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetFormData(map[string]string{
			"access_token": u.AccessToken,
			"caption":      caption,
			"image_url":    image,
		}).SetResult(&containerResponse).
		Post(containerEndpoint)

	if err != nil {
		return "", fmt.Errorf("create container req error %v", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("create container got response error %s", string(resp.Body()))
	}

	var publishResponse IGPublishResponse
	resp, err = client.R().
		SetHeaders(map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetFormData(map[string]string{
			"access_token": u.AccessToken,
			"creation_id":  containerResponse.ContainerID,
		}).SetResult(&publishResponse).
		Post(publishEndpoint)

	if err != nil {
		return "", fmt.Errorf("publish container req error %v", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("publish container got response error %s", string(resp.Body()))
	}

	return publishResponse.MediaID, nil
}

// uploadCarouselImagePost handles uploading multiple images as a carousel post to Instagram.
func (u IGUploader) uploadCarouselImagePost(images []string, caption string) (string, error) {

	client := resty.New()

	containerEndpoint := fmt.Sprintf(BaseMediaEndpoint, u.ID)
	publishEndpoint := fmt.Sprintf(MediaPublishEndpoint, u.ID)

	imageContainers := make([]string, 0)
	for _, image := range images {
		var containerResponse IGContainerResponse
		resp, err := client.R().
			SetHeaders(map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/x-www-form-urlencoded",
			}).
			SetFormData(map[string]string{
				"access_token":     u.AccessToken,
				"is_carousel_item": "true",
				"image_url":        image,
			}).SetResult(&containerResponse).
			Post(containerEndpoint)

		if err != nil {
			return "", fmt.Errorf("create container req error %v", err)
		}

		if resp.IsError() {
			return "", fmt.Errorf("create container got response error %s", string(resp.Body()))
		}

		imageContainers = append(imageContainers, containerResponse.ContainerID)
	}

	children := strings.Join(imageContainers, ",")

	var carouselResponse IGContainerResponse
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetFormData(map[string]string{
			"access_token": u.AccessToken,
			"media_type":   "CAROUSEL",
			"caption":      caption,
			"children":     children,
		}).SetResult(&carouselResponse).
		Post(containerEndpoint)

	if err != nil {
		return "", fmt.Errorf("create carousel container req error %v", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("create carousel container got response error %s", string(resp.Body()))
	}

	var publishResponse IGPublishResponse
	resp, err = client.R().
		SetHeaders(map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetFormData(map[string]string{
			"access_token": u.AccessToken,
			"creation_id":  carouselResponse.ContainerID,
		}).SetResult(&publishResponse).
		Post(publishEndpoint)

	if err != nil {
		return "", fmt.Errorf("publish carousel container req error %v", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("publish carousel container got response error %s", string(resp.Body()))
	}

	return publishResponse.MediaID, nil
}
