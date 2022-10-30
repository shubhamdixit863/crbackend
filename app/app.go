package app

import (
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"log"
	"microservicesgo/domain"
	"microservicesgo/logger"
	"microservicesgo/service"
	"net/http"
	"os"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()
	us := UserHandlers{service.NewDefaultListingService(domain.NewListingRepositoryElastic(getElasticClient())), service.NewFileUploadService(configS3())}

	// define routes
	router.
		HandleFunc("/addListing", us.addListing).
		Methods(http.MethodPost).
		Name("AddListing")

	router.
		HandleFunc("/getListing", us.getListing).
		Methods(http.MethodGet).
		Name("GetListing")

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getElasticClient() *elasticsearch.Client {

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_SERVER"),
		},
		// ...
	}
	es7, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Connected With Elastic")

	return es7

}

// configS3 creates the S3 client
func configS3() *minio.Client {

	endpoint := os.Getenv("END_POINT")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY")
	secretAccessKey := os.Getenv("AWS_SECRET_KEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}
