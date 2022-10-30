package app

import (
	"encoding/json"
	"microservicesgo/dto"
	"microservicesgo/logger"
	"microservicesgo/service"
	"net/http"
	"strings"
)

type UserHandlers struct {
	service    service.ListingService
	fileUpload service.FileUploadService
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(dto.NewCommonResponse(data)); err != nil {
		panic(err)
	}
}

func (us *UserHandlers) addListing(w http.ResponseWriter, r *http.Request) {
	files, err := HandleMultipleFileLocal(r)

	if err != nil {
		writeResponse(w, 500, "Error Uploading"+err.Error())
		return

	}
	lb, err := HandleMultiPartStringData(r)

	if err != nil {
		writeResponse(w, 500, "Error Uploading"+err.Error())
		return

	}

	go func() {
		list, err := us.fileUpload.UploadMultipleFile(files)
		if err != nil {
			logger.Error(err.Error())
			writeResponse(w, 500, "Error"+err.Error())

			return

		}
		listing, err := lb.Images(strings.Join(list[:], ",")).Build()

		if err != nil {
			logger.Error(err.Error())
			writeResponse(w, 500, "Error"+err.Error())

			return

		}

		// Call the Elastic Search Method to  Save the data

		err = us.service.AddListing(*listing)
		if err != nil {
			writeResponse(w, 500, "Error"+err.Error())
			return
		}

	}()

	writeResponse(w, 200, "SuccessFully Added Listing")
}

func (us *UserHandlers) getListing(w http.ResponseWriter, r *http.Request) {
	listing, err := us.service.GetListing()
	if err != nil {
		writeResponse(w, 200, err.Error())

		return
	}

	writeResponse(w, 200, listing)

}
