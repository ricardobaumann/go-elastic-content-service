package content

import (
	"context"
	"fmt"
	"os"

	"encoding/json"

	"github.com/olivere/elastic"

	logger "github.com/ricardo-ch/go-logger"
)

//Data ...
type repository struct {
	esCli *elastic.Client
}

// Create Repository
func NewRepository() Repository {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return repository{esCli: client}
}

func (r repository) Save(ID string, body string) (contentResponse, error) {
	result, err := r.esCli.Index().Index("contents").Type("1").BodyString(body).Id(ID).Do(context.Background())
	if err != nil {
		return contentResponse{ID: "", Body: ""}, err
	}
	logger.Info(string(result.Status))
	return contentResponse{ID: ID, Body: body}, nil
}

func (r repository) Get(ID string) (contentResponse, error) {

	//you can add your logic for database
	//r.DB.Get(&your_db_object,"your_sql_query",your_params1, ...)
	result, err := r.esCli.Get().Index("contents").Id(ID).Do(context.Background())
	if err != nil {
		return contentResponse{ID: "", Body: ""}, err
	}

	logger.Info(result.Id)

	j, _ := json.Marshal(result.Source)

	return contentResponse{ID: ID, Body: string(j)}, nil
}
