package helper

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	config "github.com/garindradeksa/socialmedia-mini/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func TypeFile(test multipart.File) bool {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	if fileType == "image/png" || fileType == "image/jpeg" {
		return true
	}
	return false
}

type Avatar struct {
	Avatar multipart.File `json:"avatar,omitempty" validate:"required"`
	// Banner multipart.File `json:"banner,omitempty" validate:"required"`
}

type Banner struct {
	Banner multipart.File `json:"banner,omitempty" validate:"required"`
}

type Content struct {
	Content multipart.File `json:"content,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

var (
	validate = validator.New()
)

type mediaUpload interface {
	AvatarUpload(avatar Avatar) (string, error)
	RemoteUpload(url Url) (string, error)
	BannerUpload(banner Banner) (string, error)
	ContentUpload(content Content) (string, error)
}

func (*media) RemoteUpload(url Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}
func (*media) BannerUpload(banner Banner) (string, error) {
	//validate
	err := validate.Struct(banner)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(banner.Banner)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) ContentUpload(content Content) (string, error) {
	//validate
	err := validate.Struct(content)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(content.Content)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) AvatarUpload(avatar Avatar) (string, error) {
	//validate
	err := validate.Struct(avatar)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(avatar.Avatar)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}
