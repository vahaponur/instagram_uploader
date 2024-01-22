# Instagram Uploader

The `instagram_uploader` package is a Go module designed for uploading images and carousels to Instagram via the Graph API. It provides a simple and efficient way to publish media on Instagram programmatically. This document outlines how to use the `instagram_uploader` package after its recent refactor.

## Features

- Upload single images with captions.
- Upload multiple images as a carousel post with captions.

## Installation

Ensure you have Go installed on your system. Install the package with the following command:

```bash
go get -u github.com/vahaponur/instagram_uploader
```

### Setting Up
Initialize the IGUploader struct with your Instagram User ID and Access Token:
```go
uploader := instagram_uploader.New("your_instagram_user_id", "your_access_token")
```

### Uploading a Single Image
To upload a single image with a caption:
```go
mediaID, err := uploader.UploadImagePost([]string{"https://example.com/path/to/image.jpg"}, "Your caption here")
if err != nil {
    log.Println(err)
}
fmt.Println("Uploaded media ID:", mediaID)
```

### Uploading a Carousel
For uploading multiple images as a carousel post with a caption:
```go
mediaID, err := uploader.UploadImagePost([]string{
    "https://example.com/path/to/image1.jpg",
    "https://example.com/path/to/image2.jpg",
    // Add more images as needed
}, "Your caption here")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Uploaded carousel media ID:", mediaID)
```

## Contributing

I am open to pull requests for new functionalities or bug fixes.
If you have ideas for improvements or have found a bug, feel free to submit a pull request or open an issue in the repository.
Your contributions are greatly appreciated!