package domain

import "microservicesgo/dto"

type Listing struct {
	Title          string `json:"title"`
	Category       string `json:"category"`
	Keywords       string `json:"keywords"`
	Location       string `json:"location"`
	Images         string `json:"files"`
	Description    string `json:"description"`
	Website        string `json:"website,omitempty"`
	Email          string `json:"email,omitempty"`
	Facebook       string `json:",omitempty"`
	Phone          string `json:",omitempty"`
	Instagram      string `json:"instagram,omitempty"`
	Linkedin       string `json:",omitempty"`
	Facilities     string `json:"facilities"`
	Pricing        string `json:"pricing"`
	OperationHours string `json:"operationHours"`
}

type ListingRepository interface {
	AddListing(listing Listing) error
	FindAllListings(size, offset int) ([]dto.ListingResponse, int, error)
	SearchListing(id string) ([]dto.ListingResponse, error)
}

// Listing builder pattern code

type ListingBuilder struct {
	listing *Listing
}

func NewListingBuilder() *ListingBuilder {
	listing := &Listing{}
	b := &ListingBuilder{listing: listing}
	return b
}

func (b *ListingBuilder) Title(title string) *ListingBuilder {
	b.listing.Title = title
	return b
}

func (b *ListingBuilder) Category(category string) *ListingBuilder {
	b.listing.Category = category
	return b
}

func (b *ListingBuilder) Keywords(keywords string) *ListingBuilder {
	b.listing.Keywords = keywords
	return b
}

func (b *ListingBuilder) Location(location string) *ListingBuilder {
	b.listing.Location = location
	return b
}

func (b *ListingBuilder) Images(images string) *ListingBuilder {
	b.listing.Images = images
	return b
}

func (b *ListingBuilder) Description(description string) *ListingBuilder {
	b.listing.Description = description
	return b
}

func (b *ListingBuilder) Website(website string) *ListingBuilder {
	b.listing.Website = website
	return b
}

func (b *ListingBuilder) Email(email string) *ListingBuilder {
	b.listing.Email = email
	return b
}

func (b *ListingBuilder) Facebook(facebook string) *ListingBuilder {
	b.listing.Facebook = facebook
	return b
}

func (b *ListingBuilder) Phone(phone string) *ListingBuilder {
	b.listing.Phone = phone
	return b
}

func (b *ListingBuilder) Instagram(instagram string) *ListingBuilder {
	b.listing.Instagram = instagram
	return b
}

func (b *ListingBuilder) Linkedin(linkedin string) *ListingBuilder {
	b.listing.Linkedin = linkedin
	return b
}

func (b *ListingBuilder) Facilities(facilities string) *ListingBuilder {
	b.listing.Facilities = facilities
	return b
}

func (b *ListingBuilder) Pricing(pricing string) *ListingBuilder {
	b.listing.Pricing = pricing
	return b
}

func (b *ListingBuilder) OperationHours(operationHours string) *ListingBuilder {
	b.listing.OperationHours = operationHours
	return b
}

func (b *ListingBuilder) Build() (*Listing, error) {
	return b.listing, nil
}
