package helper

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

func CreateImageToCloudinary(img multipart.File, c *fiber.Ctx) *uploader.UploadResult {
	cloudy, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	uploader, err := cloudy.Upload.Upload(ctx, img, uploader.UploadParams{})
	if err != nil {
		panic(err)
	}
	return uploader
}

func DeleteImageInCloudinary(imgID string, c *fiber.Ctx) {
	cloudy, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	_, err = cloudy.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: imgID,
	})
	if err != nil {
		panic(err)
	}
}
