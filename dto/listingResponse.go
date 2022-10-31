package dto

type ListingResponse struct {
	Id             string `json:"id"`
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

// ListingResponse builder pattern code

type ListingResponseBuilder struct {
	listingResponse *ListingResponse
}

func NewListingResponseBuilder() *ListingResponseBuilder {
	listingResponse := &ListingResponse{}
	b := &ListingResponseBuilder{listingResponse: listingResponse}
	return b
}

func (b *ListingResponseBuilder) Id(id string) *ListingResponseBuilder {
	b.listingResponse.Id = id
	return b
}

func (b *ListingResponseBuilder) Title(title string) *ListingResponseBuilder {
	b.listingResponse.Title = title
	return b
}

func (b *ListingResponseBuilder) Category(category string) *ListingResponseBuilder {
	b.listingResponse.Category = category
	return b
}

func (b *ListingResponseBuilder) Keywords(keywords string) *ListingResponseBuilder {
	b.listingResponse.Keywords = keywords
	return b
}

func (b *ListingResponseBuilder) Location(location string) *ListingResponseBuilder {
	b.listingResponse.Location = location
	return b
}

func (b *ListingResponseBuilder) Images(images string) *ListingResponseBuilder {
	b.listingResponse.Images = images
	return b
}

func (b *ListingResponseBuilder) Description(description string) *ListingResponseBuilder {
	b.listingResponse.Description = description
	return b
}

func (b *ListingResponseBuilder) Website(website string) *ListingResponseBuilder {
	b.listingResponse.Website = website
	return b
}

func (b *ListingResponseBuilder) Email(email string) *ListingResponseBuilder {
	b.listingResponse.Email = email
	return b
}

func (b *ListingResponseBuilder) Facebook(facebook string) *ListingResponseBuilder {
	b.listingResponse.Facebook = facebook
	return b
}

func (b *ListingResponseBuilder) Phone(phone string) *ListingResponseBuilder {
	b.listingResponse.Phone = phone
	return b
}

func (b *ListingResponseBuilder) Instagram(instagram string) *ListingResponseBuilder {
	b.listingResponse.Instagram = instagram
	return b
}

func (b *ListingResponseBuilder) Linkedin(linkedin string) *ListingResponseBuilder {
	b.listingResponse.Linkedin = linkedin
	return b
}

func (b *ListingResponseBuilder) Facilities(facilities string) *ListingResponseBuilder {
	b.listingResponse.Facilities = facilities
	return b
}

func (b *ListingResponseBuilder) Pricing(pricing string) *ListingResponseBuilder {
	b.listingResponse.Pricing = pricing
	return b
}

func (b *ListingResponseBuilder) OperationHours(operationHours string) *ListingResponseBuilder {
	b.listingResponse.OperationHours = operationHours
	return b
}

func (b *ListingResponseBuilder) Build() (*ListingResponse, error) {
	return b.listingResponse, nil
}
