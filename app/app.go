package app

import (
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/cors"
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

func Setup() *mux.Router {

	router := mux.NewRouter()
	us := UserHandlers{service.NewDefaultListingService(domain.NewListingRepositoryElastic(getElasticClient())), service.NewFileUploadService(configS3())}

	// define routes
	router.
		HandleFunc("/health", us.HealthCheck).
		Methods(http.MethodGet).
		Name("HealthCheck")
	router.
		HandleFunc("/listing", us.addListing).
		Methods(http.MethodPost).
		Name("AddListing")

	router.
		HandleFunc("/listing/search", us.getListing).
		Methods(http.MethodPost).
		Name("GetListing")

	router.
		HandleFunc("/listing/{id}", us.getListingById).
		Methods(http.MethodGet).
		Name("GetListingById")

	router.
		HandleFunc("/listing/{id}", us.deleteListingById).
		Methods(http.MethodDelete).
		Name("DeleteListingById")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	address := "localhost"
	port := "8080"
	logger.Info(fmt.Sprintf("Starting Listing  server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), handler))
	return router
}

func Start() {
	sanityCheck()
	Setup()

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
