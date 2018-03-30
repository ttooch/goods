package elastics

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"os"
	"reflect"
)

var (
	ElaClient *elastic.Client
)

func init() {

	var err error

	ElaClient, err = elastic.NewClient(
		elastic.SetURL("http://47.100.10.111:9200"),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetTraceLog(log.New(os.Stdout, "ELASTIC_Trace", log.LstdFlags)),
	)

	if err != nil {

		panic(fmt.Sprintf("Fail to init new engine: %v", err))

	}
}

type ElasticSearch interface {
	Index() string
	Type() string
	Mapping() map[string]map[string]map[string]map[string]map[string]string
	AfterFind()
	SetId(id string)
	GetSnapshot() interface{}
}

func CreateIndex(search ElasticSearch) {
	ctx := context.Background()

	_, err := ElaClient.CreateIndex(search.Index()).BodyJson(search.Mapping()).Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
}

func DelIndex(search ElasticSearch) {
	ctx := context.Background()

	_, err := ElaClient.DeleteIndex(search.Index()).Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
}

func AddDocument(search ElasticSearch) error {
	ctx := context.Background()
	_, err := ElaClient.Index().
		Index(search.Index()).
		Type(search.Type()).
		BodyJson(search).
		Do(ctx)
	return err
}

func DelDocument(search ElasticSearch, id string) error {
	ctx := context.Background()
	_, err := ElaClient.Delete().
		Index(search.Index()).
		Type(search.Type()).
		Id(id).
		Do(ctx)
	return err
}

func GetDocumentById(search ElasticSearch, id string) error {
	ctx := context.Background()
	get, err := ElaClient.Get().
		Index(search.Index()).
		Type(search.Type()).
		Id(id).
		Do(ctx)

	if !get.Found {
		return ErrNotExist
	}

	getJson, _ := get.Source.MarshalJSON()

	json.Unmarshal(getJson, search)

	search.SetId(get.Id)

	search.AfterFind()

	return err
}

func UpdateDocumentById(search ElasticSearch, id string) error {
	changed := GetChanged(search)

	ctx := context.Background()

	_, err := ElaClient.Update().Index(search.Index()).Type(search.Type()).Id(id).
		Doc(changed).
		Do(ctx)
	return err
}

func GetDocuments(search ElasticSearch, query elastic.Query, sort map[string]bool, limit ...int) (models []*ElasticSearch, err error) {

	ctx := context.Background()

	var searchResult *elastic.SearchResult

	service := ElaClient.Search().Index(search.Index()).Query(query)

	for field, value := range sort {
		service.Sort(field, value)
	}

	if len(limit) > 0 {
		models = make([]*ElasticSearch, 0, limit[0])

		searchResult, err = service.Size(limit[0]).Do(ctx)
	} else {
		models = make([]*ElasticSearch, 0)

		searchResult, err = service.Do(ctx)
	}

	if err != nil {
		return models, err
	}

	for _, hit := range searchResult.Hits.Hits {
		// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
		err := json.Unmarshal(*hit.Source, search)

		if err == nil {
			search.AfterFind()
			search.SetId(hit.Id)

			models = append(models, &search)
		} else {
			fmt.Println(err)
		}
	}

	return models, nil
}

func GetChanged(search ElasticSearch) map[string]interface{} {
	snapshot := search.GetSnapshot()

	typeSnapshot := reflect.ValueOf(snapshot)
	typeSearch := reflect.TypeOf(search)
	valueSearch := reflect.ValueOf(search)

	changed := make(map[string]interface{})

	if typeSearch.Elem().Kind() == reflect.Struct {

		for i := 0; i < typeSearch.Elem().NumField(); i++ {

			fieldName := typeSearch.Elem().Field(i).Name

			valueSnapshot := typeSnapshot.FieldByName(fieldName)

			valueSearch := valueSearch.Elem().FieldByName(fieldName)

			if valueSnapshot.Interface() != valueSearch.Interface() && fieldName != "Snapshot" {

				changed[fieldName] = valueSearch.Interface()

			}
		}
	}

	return changed
}
