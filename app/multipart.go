package app

import (
	"fmt"
	"io"
	"microservicesgo/domain"
	"microservicesgo/logger"
	"mime/multipart"
	"net/http"
	"os"
)

func HandleFileUploadLocal(r *http.Request) (*os.File, error) {

	// Parse request body as multipart form data with 32MB max memory
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// Get file uploaded via Form
	file, handler, err := r.FormFile("file")
	if err != nil {
		logger.Error("Problem In Getting File From the Form Data" + err.Error())

		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Create file locally
	dst, err := os.Create(handler.Filename)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	// Copy the uploaded file data to the newly created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return dst, nil

}

func HandleMultipleFileLocal(r *http.Request) ([]string, error) {
	var fileArray []string
	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		logger.Error(err.Error())

		return nil, err
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		filePath := "/tmp/" + files[i].Filename
		fmt.Println(files[i].Filename)
		out, err := os.Create(filePath)

		if err != nil {
			logger.Error(err.Error())

			return nil, err
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			logger.Error(err.Error())

			return nil, err
		}
		fileArray = append(fileArray, filePath)

	}
	return fileArray, nil
}

func HandleMultiPartStringData(r *http.Request) (*domain.ListingBuilder, error) {

	// Parse request body as multipart form data with 32MB max memory
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	listing := domain.NewListingBuilder().Title(r.FormValue("title")).
		Email(r.FormValue("email")).
		Category(r.FormValue("category")).
		Description(r.FormValue("description")).
		Facebook(r.FormValue("facebook")).
		Instagram(r.FormValue("instagram")).
		Facilities(r.FormValue("facilities")).
		Keywords(r.FormValue("keywords")).
		Pricing(r.FormValue("pricing")).
		OperationHours(r.FormValue("operationHours")).
		Location(r.FormValue("location"))

	return listing, nil

}
