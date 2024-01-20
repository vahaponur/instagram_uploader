package test

import (
	"github.com/spf13/viper"
	instagram_uploader "github.com/vahaponur/instagram-uploader"
	"log"
	"testing"
)

func TestCarouselPostUpload(t *testing.T) {

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
	post := instagram_uploader.CarouselImagePost{
		ImageURLS: []string{"https://cdn.metasabi.com/metasabi/product/dca1e596-94da-4cd9-a775-a0d2f6841841.jpeg", "https://cdn.metasabi.com/metasabi/product/9160fa6b-1394-4a5f-89ef-afda94132277.jpeg", "https://cdn.metasabi.com/metasabi/product/8d8202f3-1315-4879-9ba7-c6b08811258f.jpeg"},
		Caption:   "Light Yagami Portrait HQ Metal Print\n\n#metasabi #metalart",
	}
	mediaId, err := uploader.UploadCarouselImagePost(post)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf(mediaId)
}
