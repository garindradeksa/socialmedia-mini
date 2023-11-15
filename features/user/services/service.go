package services

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/garindradeksa/socialmedia-mini/config"
	"github.com/garindradeksa/socialmedia-mini/features/user"
	"github.com/garindradeksa/socialmedia-mini/helper"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}

func (uuc *userUseCase) Register(newUser user.Core) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("generate bcrypt error : ", err.Error())
		return errors.New("Unable to process password")
	}
	newUser.Password = string(hashed)
	err = uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "Username or email already exist"
		} else if strings.Contains(err.Error(), "query") {
			msg = "There is a problem with the server"
		} else {
			msg = "There is a problem with the server"
		}
		return errors.New(msg)
	}
	return nil
}

func (uuc *userUseCase) Login(username, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(username)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Data not found"
		} else {
			msg = "There is a problem with the server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := helper.CheckPassword(res.Password, password); err != nil {
		log.Println("Failed to compare password : ", err.Error())
		return "", user.Core{}, errors.New("Email or password doesn't match")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))

	return useToken, res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("user not found")
	}

	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Data not found"
		} else {
			msg = "There is a problem with the server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(formHeader multipart.FileHeader, formHeader2 multipart.FileHeader, token interface{}, updatedProfile user.Core) (user.Core, error) {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return user.Core{}, errors.New("Data not found")
	}

	updatedProfile.ID = uint(id)

	if formHeader.Size > 5000000 {
		return user.Core{}, errors.New("file size is too big")
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return user.Core{}, errors.New("open formheader error")
	}

	if !helper.TypeFile(formFile) {
		return user.Core{}, errors.New("use jpg or png type file")
	}
	defer formFile.Close()
	formFile, _ = formHeader.Open()
	uploadUrl, err := helper.NewMediaUpload().AvatarUpload(helper.Avatar{Avatar: formFile})

	if err != nil {
		return user.Core{}, errors.New("server error")
	}

	updatedProfile.Avatar = uploadUrl

	if formHeader2.Size > 5000000 {
		return user.Core{}, errors.New("file size is too big")
	}

	formFile2, err := formHeader2.Open()
	if err != nil {
		return user.Core{}, errors.New("open formheader error")
	}

	if !helper.TypeFile(formFile2) {
		return user.Core{}, errors.New("use jpg or png type file")
	}
	defer formFile2.Close()
	formFile2, _ = formHeader2.Open()
	uploadUrl2, err := helper.NewMediaUpload().BannerUpload(helper.Banner{Banner: formFile2})

	if err != nil {
		return user.Core{}, errors.New("server error")
	}
	updatedProfile.Banner = uploadUrl2

	res, err := uuc.qry.Update(updatedProfile)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Failed to update, no new record or data not found"
		} else {
			msg = "There is a problem with the server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Deactivate(token interface{}) error {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return errors.New("Data not found")
	}

	err := uuc.qry.Deactivate(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Data not found"
		} else {
			msg = "There is a problem with the server"
		}
		return errors.New(msg)
	}

	return nil
}

// Done
