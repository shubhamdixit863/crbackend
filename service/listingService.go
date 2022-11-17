package service

import (
	"microservicesgo/domain"
	"microservicesgo/dto"
)

type ListingService interface {
	AddListing(listing domain.Listing) error
	GetListing(size, offset int) (interface{}, error)
	SearchListing(id string) ([]dto.ListingResponse, error)
}

type DefaultListingService struct {
	repo domain.ListingRepository
}

func (dls *DefaultListingService) AddListing(listing domain.Listing) error {

	return dls.repo.AddListing(listing)

}

func (dls *DefaultListingService) GetListing(size, offset int) (interface{}, error) {
	listing, total, err := dls.repo.FindAllListings(size, offset)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]interface{})
	mp["listing"] = listing
	mp["total"] = total
	return mp, nil
}
func (dls *DefaultListingService) SearchListing(id string) ([]dto.ListingResponse, error) {

	return dls.repo.SearchListing(id)

}

func NewDefaultListingService(repo domain.ListingRepository) *DefaultListingService {

	return &DefaultListingService{
		repo: repo,
	}

}
