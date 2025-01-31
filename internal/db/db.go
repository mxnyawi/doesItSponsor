package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/joho/godotenv"
)

// Define a struct to hold your database connection information
type Database struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
}

// Initialize database connection
func NewDatabase() (*Database, error) {
	// Load environment variables from config.env file
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading config.env file: %v", err)
		return nil, err
	}

	// Update this to your cluster details
	connectionString := os.Getenv("CAPELLA_CONNECTION_STRING")
	username := os.Getenv("CAPELLA_USERNAME")
	password := os.Getenv("CAPELLA_PASSWORD")

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	}

	// Sets a pre-configured profile called "wan-development" to help avoid latency issues
	// when accessing Capella from a different Wide Area Network
	// or Availability Zone (e.g. your laptop).
	if err := options.ApplyProfile(gocb.
		ClusterConfigProfileWanDevelopment); err != nil {
		log.Fatal(err)
	}

	// Initialize the Connection
	cluster, err := gocb.Connect(connectionString, options)
	if err != nil {
		log.Fatal(err)
	}

	// Open the default bucket
	bucket := cluster.Bucket("sponsorcsv")
	if err := bucket.WaitUntilReady(10*time.Second, nil); err != nil {
		log.Fatalf("Failed to wait for bucket readiness: %v", err)
		return nil, err
	}

	return &Database{
		Cluster: cluster,
		Bucket:  bucket,
	}, nil
}

// Close database connection
func (db *Database) Close() {
	db.Cluster.Close(nil)
}

func BuildQuery(target, field string) string {
	return fmt.Sprintf("SELECT * FROM sponsorcsv._default._default  WHERE LOWER(%s) LIKE  LOWER('%%%s%%')", field, target)
}

type Organisation struct {
	OrganisationName string      `json:"organisation_name"`
	City             string      `json:"city"`
	County           interface{} `json:"county"`
	Type             string      `json:"type"`
	Route            string      `json:"route"`
}

// Define a function to fetch a document from the Database
func (db *Database) GetDocument(query string) ([]*Organisation, error) {
	rows, err := db.Cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	organisations, err := db.ParseRows(rows)
	if err != nil {
		return nil, err
	}

	return organisations, nil
}

// Define a function to parse the rows returned from the Database
func (db *Database) ParseRows(rows *gocb.QueryResult) ([]*Organisation, error) {
	var organisations []*Organisation
	for rows.Next() {

		var result map[string]Organisation
		if err := rows.Row(&result); err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}

		for _, organisation := range result {
			organisations = append(organisations, &organisation)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return organisations, nil
}
