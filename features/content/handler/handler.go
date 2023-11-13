package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/garindradeksa/socialmedia-mini/dtos"
	"github.com/garindradeksa/socialmedia-mini/features/content"
	"github.com/garindradeksa/socialmedia-mini/helper"

	"github.com/labstack/echo/v4"
)

type contentHandle struct {
	srv content.ContentService
}

func New(cs content.ContentService) content.ContentHandler {
	return &contentHandle{
		srv: cs,
	}
}

func (ch *contentHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {

		formHeader, err := c.FormFile("content")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}

		input := AddContentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}

		cnv := ToCore(input)

		res, err := ch.srv.Add(*formHeader, c.Get("user"), *cnv)
		if err != nil {
			log.Println("error post content : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success post content", ToAddResponse(res)))
	}
}

func (ch *contentHandle) ContentDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		contentID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}
		res, err := ch.srv.ContentDetail(uint(contentID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success get user content", res))
	}
}

func (ch *contentHandle) ContentList() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ch.srv.ContentList()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success get all content", ListCoreToResp(res)))
	}
}

func (ch *contentHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")

		contentID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		body := UpdateContentRequest{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := ch.srv.Update(token, uint(contentID), *ToCore(body))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Success edit content", res))
	}
}

func (ch *contentHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")

		contentID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		err = ch.srv.Delete(token, uint(contentID))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "Success delete content")
	}
}

func (ch *contentHandle) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramUsername := c.Param("username")

		res, err := ch.srv.GetProfile(paramUsername)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success get user content", res))
	}
}
