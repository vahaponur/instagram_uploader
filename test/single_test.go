package test

import (
	"github.com/spf13/viper"
	instagram_uploader "github.com/vahaponur/instagram-uploader"
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
	}
	post := instagram_uploader.SinglePost{
		ImageURL: "https://cdn.metasabi.com/product/c23fb672-9017-4185-84c3-f6db3c63e7b8.jpeg",
		Caption:  "#metasabi #metalart",
	}
	mediaId, err := uploader.UploadSinglePost(post)
	if err != nil {
		t.Error(err)
	}
	t.Logf(mediaId)
}