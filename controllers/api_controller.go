package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"test-app/utils"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

type APIController struct{}

func (*APIController) TransformImage(c *gin.Context) {

	errorUtils := utils.ErrorUtils{}

	// Getting the file from request
	file, err := c.FormFile("image")
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "Please provide a valid image file of format .jpeg or .png", err)
		return
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusInternalServerError, "Can't open the provided file. Please try again", err)
		return
	}
	defer f.Close()

	// Extracting all the queries
	var queries utils.TransformImageQuery
	if err := c.ShouldBindQuery(&queries); err != nil {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "Please provide valid queries for your specific transformation needs", err)
		return
	}

	// Validate queries
	if ok, err := queries.Validate(); !ok {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "", err)
		return
	}

	// Validate image
	// For now, just leave it empty.

	// Decode the file to image
	imageData, err := imaging.Decode(f)
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusInternalServerError, "Error occured while decoding provided image", err)
		return
	}

	// Transform image from the provided query
	imageUtils := utils.ImageUtils{}
	transformedBytes, _ := imageUtils.TransformImage(&imageData, &queries)

	c.Data(http.StatusOK, "image/jpeg", transformedBytes)
}

func (*APIController) ApplyFilter(c *gin.Context) {

	errorUtils := utils.ErrorUtils{}

	// Getting the file from request
	file, err := c.FormFile("image")
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "Please provide a valid image file of format .jpeg or .png", err)
		return
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusInternalServerError, "Can't open the provided file. Please try again", err)
		return
	}
	defer f.Close()

	// Getting query from url
	filterName := c.Query("name")
	fmt.Println("Given filter name: ", filterName)
	if filterName == "" {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "Please provide a valid filter name. It should be either `grayscale`, `sepia` or `negative`", errors.New(""))
		return
	}

	if filterName == "grayscale" || filterName != "sepia" || filterName == "negative" {
		fmt.Println("Filter found!!")
	} else {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "Please provide a valid filter name. It should be either `grayscale`, `sepia` or `negative`", errors.New(""))
		return
	}

	// Decode the file to image
	imageData, err := imaging.Decode(f)
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusInternalServerError, "Error occured while decoding provided image", err)
		return
	}

	imageUtils := utils.ImageUtils{}

	var filteredImage []byte

	if filterName == "grayscale" {
		filteredImage, err = imageUtils.ApplyCustomFilter(&imageData, imageUtils.Grayscale)
	} else if filterName == "sepia" {
		filteredImage, err = imageUtils.ApplyCustomFilter(&imageData, imageUtils.Sepia)
	} else if filterName == "negative" {
		filteredImage, err = imageUtils.ApplyCustomFilter(&imageData, imageUtils.Negative)
	} else {
		err = errors.New("error occured while applying filter")
	}

	if err != nil {
		errorUtils.SendJSONError(c, http.StatusInternalServerError, "Error occured while applying filter to the provided image", err)
		return
	}

	c.Data(http.StatusOK, "image/jpeg", filteredImage)
}

func (*APIController) Meta(c *gin.Context) {
	errorUtils := utils.ErrorUtils{}

	// Getting the file from request
	file, err := c.FormFile("image")
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusBadRequest, "Please provide a valid image file of format .jpeg or .png", err)
		return
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		errorUtils.SendJSONError(c, http.StatusInternalServerError, "Can't open the provided file. Please try again", err)
		return
	}
	defer f.Close()

	imageUtils := utils.ImageUtils{}
	meta := imageUtils.GetMetaData(&f)

	c.JSON(http.StatusOK, meta)
}
