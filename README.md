# Instagram Uploader

The `instagram_uploader` package is a Go module designed for uploading images and carousels to Instagram via the Graph API. 
It provides a simple and efficient way to publish media on Instagram programmatically. 
This document outlines how to use the `instagram_uploader` package.

## Features

- Upload single images with captions.
- Upload multiple images as a carousel post with captions.

## Installation

To use the `instagram_uploader` package, first, ensure you have Go installed on your system. Then, you can install the package using the following command:

```bash
go get -u github.com/vahaponur/instagram_uploader
```
## Usage

### Importing the Package

Import the `instagram_uploader` package into your Go project:

```go
import "github.com/vahaponur/instagram_uploader"
```
## Setting Up
### Initialize the IGUploader struct with your Instagram User ID, Access Token, and the API version you want to use:
```go
uploader := instagram_uploader.IGUploader{
    ID:          "your_instagram_user_id",
    AccessToken: "your_access_token",
    Version:     "v18.0",
}
```
## Uploading a Single Image
### To upload a single image, create a SingleImagePost struct and use the UploadSingleImagePost method:
```go
post := instagram_uploader.SingleImagePost{
    ImageURL: "https://example.com/path/to/image.jpg",
    Caption:  "Your caption here",
}

mediaID, err := uploader.UploadSingleImagePost(post)
if err != nil {
    log.Println(err)
}
fmt.Println("Uploaded media ID:", mediaID)
```
## Uploading a Carousel
### For uploading a carousel, create a CarouselImagePost struct and use the UploadCarouselImagePost method:
```go
carouselPost := instagram_uploader.CarouselImagePost{
    ImageURLS: []string{
        "https://example.com/path/to/image1.jpg",
        "https://example.com/path/to/image2.jpg",
        // Add more images as needed
    },
    Caption: "Your caption here",
}

mediaID, err := uploader.UploadCarouselImagePost(carouselPost)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Uploaded carousel media ID:", mediaID)
```
## Contributing

I am open to pull requests for new functionalities or bug fixes.
If you have ideas for improvements or have found a bug, feel free to submit a pull request or open an issue in the repository. 
Your contributions are greatly appreciated!


