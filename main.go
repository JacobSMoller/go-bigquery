import (
	"log"
	"cloud.google.com/go/bigquery"
)

c, err := bigquery.NewClient(ctx, "dev-vmlds-cm")
if err != nil {
	log.Fatal(err.Error())
}

