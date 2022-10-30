package domain

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"microservicesgo/logger"
)

type ListingRepositoryElastic struct {
	client *elasticsearch.Client
}

func (l ListingRepositoryElastic) AddListing(listing Listing) error {

	res, err := l.client.Index("listing", esutil.NewJSONReader(&listing))
	logger.Info(res.String())
	if err != nil {
		return err
	}

	return nil

}

func (l ListingRepositoryElastic) FindAllListings() ([]Listing, error) {
	var r map[string]interface{}

	var listing []Listing

	res, err := l.client.Search(l.client.Search.WithIndex("listing"), l.client.Search.WithPretty())

	if err != nil {
		return listing, err
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return listing, err
	}

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		sourceMap := source.(map[string]interface{})
		//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"])

		li, _ := NewListingBuilder().Location(sourceMap["location"].(string)).
			Title(sourceMap["title"].(string)).
			Images(sourceMap["files"].(string)).
			Location(sourceMap["location"].(string)).
			Description(sourceMap["description"].(string)).
			Pricing(sourceMap["pricing"].(string)).
			Facilities(sourceMap["facilities"].(string)).
			Build()

		listing = append(listing, *li)

	}

	return listing, nil

}

func NewListingRepositoryElastic(client *elasticsearch.Client) *ListingRepositoryElastic {
	return &ListingRepositoryElastic{
		client: client,
	}
}
