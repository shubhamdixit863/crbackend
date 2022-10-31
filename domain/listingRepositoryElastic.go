package domain

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"microservicesgo/dto"
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

func (l ListingRepositoryElastic) FindAllListings() ([]dto.ListingResponse, error) {
	var r map[string]interface{}

	var listing []dto.ListingResponse

	res, err := l.client.Search(l.client.Search.WithIndex("listing"), l.client.Search.WithPretty())

	if err != nil {
		return listing, err
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return listing, err
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		sourceMap := source.(map[string]interface{})

		li, _ := dto.NewListingResponseBuilder().Location(sourceMap["location"].(string)).
			Title(sourceMap["title"].(string)).
			Images(sourceMap["files"].(string)).
			Location(sourceMap["location"].(string)).
			Description(sourceMap["description"].(string)).
			Pricing(sourceMap["pricing"].(string)).
			Facilities(sourceMap["facilities"].(string)).
			Id(hit.(map[string]interface{})["_id"].(string)).
			Keywords(sourceMap["keywords"].(string)).
			Build()

		listing = append(listing, *li)

	}

	return listing, nil

}

func (l ListingRepositoryElastic) SearchListing(id string) ([]dto.ListingResponse, error) {
	var r map[string]interface{}
	var listing []dto.ListingResponse

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"terms": map[string]interface{}{
				"_id": []string{id},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return listing, err
	}

	res, err := l.client.Search(l.client.Search.WithIndex("listing"), l.client.Search.WithPretty(), l.client.Search.WithBody(&buf))

	if err != nil {
		return listing, err
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return listing, err
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		sourceMap := source.(map[string]interface{})

		li, _ := dto.NewListingResponseBuilder().Location(sourceMap["location"].(string)).
			Title(sourceMap["title"].(string)).
			Images(sourceMap["files"].(string)).
			Location(sourceMap["location"].(string)).
			Description(sourceMap["description"].(string)).
			Pricing(sourceMap["pricing"].(string)).
			Facilities(sourceMap["facilities"].(string)).
			Id(hit.(map[string]interface{})["_id"].(string)).
			Keywords(sourceMap["keywords"].(string)).
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
