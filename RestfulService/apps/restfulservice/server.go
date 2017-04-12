package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	//"strconv"
)

func main() {

	// todo step 1 : open db
	db, error := gorm.Open("sqlite3", "test.db")
	if error != nil {
		log.Fatal(error)
		panic(error)
	}
	db.LogMode(true)

	// todo step 2 : migrate Thing as model updates
	db.AutoMigrate(&Thing{})

	// context reference to db if open / create success
	c := ThingContext{DB: db}

	// Router configuration
	r := mux.NewRouter()
	r.HandleFunc("/", c.HandleHome)
	r.HandleFunc("/thing/{id}", c.HandleGetThing).Methods("GET")
	r.HandleFunc("/thing", c.HandleCreateThing).Methods("POST")
	r.HandleFunc("/things/{id}", c.HandleUpdateThing).Methods("POST")
	r.HandleFunc("/things/{id}", c.HandleDeleteThing).Methods("DELETE")
	r.HandleFunc("/things", c.HandleGetThings).Methods("GET")

	// Router listener
	log.Fatal(http.ListenAndServe(":8088", r))
}

type ThingContext struct {
	DB *gorm.DB
}

func (c *ThingContext) HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home")
}

func (c *ThingContext) HandleGetThing(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	log.Printf("Retrieving Thing with id : %s\n", vars["id"])

	thing := Thing{}

	if err := c.DB.First(&thing, vars["id"]).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	data, error := json.Marshal(thing)
	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

func (c *ThingContext) HandleCreateThing(writer http.ResponseWriter, request *http.Request) {
	// step 1 : create Thing reference object
	thing := Thing{}

	// step 2 : decode data from request body
	decoder := json.NewDecoder(request.Body)

	// step 3 : decode decoder value to Thing type
	err := decoder.Decode(&thing)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// step 4 : create thing object entry in db
	if err := c.DB.Create(&thing).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	// parse response
	data, error := json.Marshal(thing)
	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	// Write Response
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

func (c *ThingContext) HandleUpdateThing(writer http.ResponseWriter, request *http.Request) {
	// step 1 : get parameters from url path
	vars := mux.Vars(request)

	// step 2 : create the reference object and get-assign data from table
	thing := Thing{}
	if err := c.DB.First(&thing, vars["id"]).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	// step 3: Decode POST body in different reference object
	updateThing := Thing{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&updateThing)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// step 4 : update decoded value into reference object
	thing.Name = updateThing.Name
	thing.Description = updateThing.Description

	// step 5: save updated value to db
	if err := c.DB.Save(&thing).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	// step 6 : parse response
	data, err := json.Marshal(thing)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// step 7 : write response
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

func (c *ThingContext) HandleDeleteThing(writer http.ResponseWriter, request *http.Request) {
	// step 1 : get request params
	vars := mux.Vars(request)

	// step 2 : create reference object and get-assign data from table
	thing := Thing{}
	if err := c.DB.First(&thing, vars["id"]).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := c.DB.Delete(&thing).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	writer.WriteHeader(http.StatusOK)
}

func (c *ThingContext) HandleGetThings(writer http.ResponseWriter, request *http.Request) {

	//limit := 10
	//offset := 0
	//
	//if len(request.URL.Query()["limit"]) > 0 {
	//	l, err := strconv.Atoi(request.URL.Query()["limit"][0])
	//	if err != nil {
	//		log.Fatal(err)
	//		log.Printf("Invalid limit")
	//		panic(err)
	//	}
	//	limit = l
	//}
	//
	//if len(request.URL.Query()["offset"]) > 0 {
	//	o, err := strconv.Atoi(request.URL.Query()["offset"][0])
	//	if err != nil {
	//		log.Fatal(err)
	//		log.Printf("Invalid offset")
	//		panic(err)
	//	}
	//	offset = o
	//}

	// step 1 : create things object
	things := Things{}

	// step 2 : get things and limit to 10
	if err := c.DB.Limit(10).Find(&things).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("things", things)

	// step 3 : parse response
	data, err := json.Marshal(things)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	//step 4 : write response
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}
