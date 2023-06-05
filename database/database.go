package database

import (
	"TlacuachesAsesinos/constants"
	"TlacuachesAsesinos/model"
	"context"
	"fmt"
	_ "io/ioutil"
	"log"
	_ "strconv"
	"time"

	//"golang.org/x/net/context"
	//"golang.org/x/oauth2/google"
	//"gopkg.in/Iwark/spreadsheet.v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Criteria struct {
	Field       string
	Restriction string
	Value       interface{}
}
type MongoLocal struct{}

func CheckCollectionsExist() (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constants.Mongo_uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db := client.Database(constants.Mongo_database)
	if err != nil {
		// Handle error
		log.Printf("Failed to get coll names: %v", err)
		return nil, err
	}
	for _, collection := range constants.Mongo_CollectionNames {
		//https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.11.2/mongo#Database.CreateCollection
		db.CreateCollection(context.TODO(), collection)
	}
	return true, nil
}

func Connect(criteria bson.M, collection string, fn func(ctx context.Context, criteria bson.M) (interface{}, error)) (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constants.Mongo_uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//coll := client.Database(dt.Database_Name).Collection(collection)
	output, err := fn(context.TODO(), criteria)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return output, nil
}

func checkError(err error, errString string) {
	if err != nil {
		log.Println(errString)
		panic(err.Error())
	}
}

func Search(folio, bot string) (rg model.Registro) {
	res, err := findOne(BuildCriteria(BasicCriteria(bot, folio)), constants.Mongo_collection)
	if err != nil {
		return
	}
	err = res.Decode(&rg)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func SearchPage(offset int, bot string) (prev int, next int, buttons [][]string) {
	prev = offset - 1
	next = offset + 1
	folios := []model.Registro{}
	offset *= constants.Const_pages

	ops := options.Find().SetSort(bson.D{{"folio", 1}}).SetSkip(int64(offset)).SetLimit(int64(constants.Const_pages))
	res, total, err := findMany(BuildCriteria(BasicCriteria(bot, "")), constants.Mongo_collection, ops)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	if res == nil {
		return
	}
	if err = res.All(context.TODO(), &folios); err != nil {
		return //panic(err)
	}

	var curr = int64(len(folios) + offset)
	if total == curr {
		next = -1
	}
	for _, rg := range folios {
		buttons = append(buttons, []string{fmt.Sprintf("%v - %v", rg.Folio, rg.Nombre), fmt.Sprint("Folio - ", rg.Folio)})
	}
	return
}

func Save(rg model.Registro, pantalla int) (string, string) {
	dateS, dateE, clockS, clockE, numeroExterior := "", "", "", "", ""
	var err error
	var res interface{}
	yS, mS, dS := rg.FechaSalida.Date()
	hS, minS, sS := rg.HoraSalida.Clock()

	yE, mE, dE := rg.FechaEntrada.Date()
	hE, minE, sE := rg.HoraEntrada.Clock()

	if !rg.FechaSalida.IsZero() {
		dateS = fmt.Sprintf("%02d/%02d/%04d", dS, mS, yS)
	}
	if !rg.HoraSalida.IsZero() {
		clockS = fmt.Sprintf("%02d:%02d:%02d", hS, minS, sS)
	}
	if !rg.FechaEntrada.IsZero() {
		dateE = fmt.Sprintf("%02d/%02d/%04d", dE, mE, yE)
	}
	if !rg.HoraEntrada.IsZero() {
		clockE = fmt.Sprintf("%02d:%02d:%02d", hE, minE, sE)
	}
	if rg.NumeroExterior != 0 {
		numeroExterior = fmt.Sprint(rg.NumeroExterior)
	}

	if rg.Folio == "" {
		rg.Folio = fmt.Sprint(time.Now().Unix())
	}
	if pantalla == 0 {
		rg.Creacion = time.Now()
		rg.Estatus = constants.Const_estatus_por_entrar
		res, err = save(rg)
	} else {
		rg.Estatus = constants.Const_estatus_completo
		if rg.FechaSalida.IsZero() && rg.HoraSalida.IsZero() {
			rg.Estatus = constants.Const_estatus_por_salir

		}
		res, err = update(rg)
	}
	checkError(err, "Error on Update/Save")
	fmt.Println(dateS, dateE, clockS, clockE, numeroExterior, res)
	return fmt.Sprintf("%v - %v", rg.Folio, rg.Nombre), rg.Folio
}

func save(insert interface{}) (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constants.Mongo_uri))
	if err != nil {
		log.Println("database.go:192 ", err)
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin insertOne
	coll := client.Database(constants.Mongo_database).Collection(constants.Mongo_collection)
	result, err := coll.InsertOne(context.TODO(), insert)
	if err != nil {
		//panic(err)
		log.Println("database.go:205 ", err)
		return nil, err
	}

	log.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return result, nil
}

func update(insert model.Registro) (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constants.Mongo_uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin insertOne
	coll := client.Database(constants.Mongo_database).Collection(constants.Mongo_collection)
	result, err := coll.ReplaceOne(context.TODO(), bson.M{"folio": insert.Folio}, insert)
	if err != nil {
		//panic(err)
		log.Println("database.go:216 ", err)
		return nil, err
	}

	log.Printf("Document inserted with ID: %s\n", result.UpsertedID)
	return result, nil
}

func findOne(criteria bson.M, collection string, opt ...*options.FindOptions) (*mongo.SingleResult, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constants.Mongo_uri))
	if err != nil {
		log.Println("database.go:238 ", err)
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(constants.Mongo_database).Collection(collection)
	output := coll.FindOne(context.TODO(), criteria)
	if err != nil {
		log.Println("database.go:249 ", err)
		return nil, err
	}
	return output, nil
}
func findMany(criteria bson.M, collection string, opt ...*options.FindOptions) (*mongo.Cursor, int64, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constants.Mongo_uri))
	if err != nil {
		log.Println("database.go:238 ", err)
		return nil, 0, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(constants.Mongo_database).Collection(collection)
	count, err := coll.CountDocuments(context.TODO(), criteria)
	if err != nil {
		panic(err)
	}
	output, err := coll.Find(context.TODO(), criteria, opt[0])
	if err != nil {
		log.Println("database.go:249 ", err)
		return nil, 0, err
	}
	return output, count, nil
}

func BuildCriteria(criteria []Criteria) (multi bson.M) {
	//var one interface{}
	multi = bson.M{}
	for _, ctr := range criteria {
		if ctr.Restriction == "" {
			multi[ctr.Field] = ctr.Value
		} else {
			multi[ctr.Field] = bson.D{{ctr.Restriction, ctr.Value}}
		}
	}
	return
}
func BasicCriteria(bot, folio string) (criteria []Criteria) {
	criteria = append(criteria, Criteria{Value: bot, Field: "bot"})
	if folio != "" {
		criteria = append(criteria, Criteria{Value: folio, Field: "folio"})
	}
	return
}
