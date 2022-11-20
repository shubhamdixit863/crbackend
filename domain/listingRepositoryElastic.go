package domain

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	esutils "github.com/olivere/elastic/v7"
	"io"
	"log"
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

func (l ListingRepositoryElastic) FindAllListings(size, offset int, search, category, location string) ([]dto.ListingResponse, int, error) {

	var r map[string]interface{}
	var total int

	var listing []dto.ListingResponse

	if location != "" || category != "" || search != "" {
		q := esutils.NewBoolQuery()
		q = q.Should(esutils.NewMatchQuery("location", location), esutils.NewMatchQuery("category", category), esutils.NewMatchQuery("description", search), esutils.NewMatchQuery("title", search))

		src, err := q.Source()

		var buf bytes.Buffer
		query := map[string]interface{}{
			"query": src,
		}
		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query: %s", err)
		}

		res, err := l.client.Search(l.client.Search.WithIndex("listing"),
			l.client.Search.WithPretty(),
			l.client.Search.WithSize(size),
			l.client.Search.WithFrom(offset),
			l.client.Search.WithTrackTotalHits(true),
			l.client.Search.WithBody(&buf),
		)

		if err != nil {
			logger.Error(err.Error())
			return listing, total, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)

		if res.IsError() {
			var e map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				logger.Error("Error parsing the response body: %s" + err.Error())
				return listing, total, err

			} else {
				// Print the response status and error information.
				logger.Error("Error parsing the response body: %s" + e["error"].(map[string]interface{})["type"].(string) + e["error"].(map[string]interface{})["reason"].(string))

				return listing, total, err

			}
		}

		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			logger.Error(err.Error())

			return listing, total, err
		}

		total = int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

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

		return listing, total, err

	} else {
		res, err := l.client.Search(l.client.Search.WithIndex("listing"),
			l.client.Search.WithPretty(),
			l.client.Search.WithSize(size),
			l.client.Search.WithFrom(offset),
			l.client.Search.WithTrackTotalHits(true),
		)

		if err != nil {
			logger.Error(err.Error())
			return listing, total, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)

		if res.IsError() {
			var e map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				logger.Error("Error parsing the response body: %s" + err.Error())
				return listing, total, err

			} else {
				// Print the response status and error information.
				logger.Error("Error parsing the response body: %s" + e["error"].(map[string]interface{})["type"].(string) + e["error"].(map[string]interface{})["reason"].(string))

				return listing, total, err

			}
		}

		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			logger.Error(err.Error())

			return listing, total, err
		}

		total = int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

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

		return listing, total, err

	}

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
