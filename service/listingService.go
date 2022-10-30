package service

import "microservicesgo/domain"

type ListingService interface {
	AddListing(listing domain.Listing) error
	GetListing() ([]domain.Listing, error)
}

type DefaultListingService struct {
	repo domain.ListingRepository
}

func (dls *DefaultListingService) AddListing(listing domain.Listing) error {

	return dls.repo.AddListing(listing)

}

func (dls *DefaultListingService) GetListing() ([]domain.Listing, error) {

	return dls.repo.FindAllListings()

}

func NewDefaultListingService(repo domain.ListingRepository) *DefaultListingService {

	return &DefaultListingService{
		repo: repo,
	}

}
