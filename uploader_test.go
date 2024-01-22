package instagram_uploader

import "testing"

// Test data
var sampleImages = []string{
	"https://cdn.metasabi.com/metasabi/product/dca1e596-94da-4cd9-a775-a0d2f6841841.jpeg",
	"https://cdn.metasabi.com/metasabi/product/9160fa6b-1394-4a5f-89ef-afda94132277.jpeg",
	"https://cdn.metasabi.com/metasabi/product/8d8202f3-1315-4879-9ba7-c6b08811258f.jpeg",
}

func TestUploadSingleImage(t *testing.T) {

	uploader := New("ID", "TOKEN")
	mediaID, err := uploader.UploadImagePost([]string{sampleImages[0]}, "Test image")
	if err != nil {
		t.Errorf("Failed to upload single image: %v", err)
	}
	if mediaID == "" {
		t.Error("Expected a media ID for single image upload, got empty string")
	}
}

func TestUploadMultiImage(t *testing.T) {

	uploader := New("ID", "TOKEN")
	mediaID, err := uploader.UploadImagePost(sampleImages, "Test image")
	if err != nil {
		t.Errorf("Failed to upload multi image: %v", err)
	}
	if mediaID == "" {
		t.Error("Expected a media ID for multi image upload, got empty string")
	}
}
