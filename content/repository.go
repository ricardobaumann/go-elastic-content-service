package content

import (
	"context"
	"fmt"
	"os"

	"github.com/olivere/elastic"
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

func (r repository) Get(ID string) (contentResponse, error) {

	//you can add your logic for database
	//r.DB.Get(&your_db_object,"your_sql_query",your_params1, ...)
	result, err := r.esCli.Get().Index("contents").Id(ID).Do(context.Background())
	if err != nil {
		fmt.Print("error")
	}

	fmt.Print(result.Source)

	return contentResponse{ID}, nil
}
