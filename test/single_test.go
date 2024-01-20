package test

import (
	"github.com/spf13/viper"
	"github.com/vahaponur/instagram_uploader"
	"log"
	"testing"
)

func TestSinglePostUpload(t *testing.T) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
	id := viper.GetString("id")
	token := viper.GetString("token")
	uploader := instagram_uploader.IGUploader{
		ID:          id,
		AccessToken: token,
		Version:     "v18.0",
	}
	post := instagram_uploader.SingleImagePost{
		ImageURL: "https://cdn.metasabi.com/product/c23fb672-9017-4185-84c3-f6db3c63e7b8.jpeg",
		Caption:  "#metasabi #metalart",
	}
	mediaId, err := uploader.UploadSingleImagePost(post)
	if err != nil {
		t.Error(err)
	}
	t.Logf(mediaId)
}
