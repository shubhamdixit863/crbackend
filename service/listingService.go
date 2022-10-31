package service

import (
	"microservicesgo/domain"
	"microservicesgo/dto"
)

type ListingService interface {
	AddListing(listing domain.Listing) error
	GetListing() ([]dto.ListingResponse, error)
	SearchListing(id string) ([]dto.ListingResponse, error)
}

type DefaultListingService struct {
	repo domain.ListingRepository
}

func (dls *DefaultListingService) AddListing(listing domain.Listing) error {

	return dls.repo.AddListing(listing)

}

func (dls *DefaultListingService) GetListing() ([]dto.ListingResponse, error) {

	return dls.repo.FindAllListings()

}
func (dls *DefaultListingService) SearchListing(id string) ([]dto.ListingResponse, error) {

	return dls.repo.SearchListing(id)

}

func NewDefaultListingService(repo domain.ListingRepository) *DefaultListingService {

	return &DefaultListingService{
		repo: repo,
	}

}
