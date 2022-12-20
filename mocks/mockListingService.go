package mocks

import (
	"microservicesgo/domain"
	"microservicesgo/dto"
)

type MockListingService struct {
}

func (mls *MockListingService) AddListing(listing domain.Listing) error {
	return nil

}
func (mls *MockListingService) GetListing(request *dto.SearchRequest) (interface{}, error) {
	return nil, nil

}
func (mls *MockListingService) SearchListing(id string) ([]dto.ListingResponse, error) {
	return nil, nil

}
func (mls *MockListingService) DeleteListing(id string) error {
	return nil

}