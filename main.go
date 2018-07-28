package main

import (
	"os"
	"gokit-practice/abilities"
	"github.com/joho/godotenv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	//"github.com/satori/go.uuid"
	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	env := GetEnv()

	db, _ := gorm.Open(
		"postgres",
		"host="+env.dbHost+" port="+env.dbPort+" user="+env.dbUser+" dbname="+env.dbName+" sslmode=disable",
	)
	defer db.Close()

	abilitiesService := abilities.NewService(db)
	r := abilities.MakeHTTPHandler(abilitiesService, logger)

	spew.Dump("Starting server")

	err := http.ListenAndServe(":3000", r)

	spew.Dump(err)

	//// test
	//abilitiesService := abilities.NewService(db)
	//createAbility := abilities.MakeCreateAbilityEndpoint(abilitiesService)
	//abilityID := uuid.NewV4()
	//ownerID := uuid.NewV4() //uuid.FromString("3eac1204-da1d-42a3-8b24-08e884fbe72e")
	//ability := abilities.Ability{
	//	ID: abilityID,
	//	OwnerId: ownerID,
	//	Caption: "asdf",
	//}
	//abilityRequest := abilities.CreateAbilityRequest{
	//	Ability: ability,
	//}
	//a, err := createAbility(nil, abilityRequest)
	//spew.Dump(a, err)


	//var a []abilities.Ability
	////db.Find(&a)
	//db.Where("owner_id = ?", "3eac1204-da1d-42a3-8b24-08e884fbe72e").Find(&a)
	//spew.Dump(a)


	//getAbility := abilities.MakeGetAbilityEndpoint(abilitiesService)
	//abilityID, _ := uuid.FromString("569e699f-a98d-4651-b3e9-c27ea308a518")
	//ownerID, _ := uuid.FromString("3eac1204-da1d-42a3-8b24-08e884fbe72e")
	//
	//ability := abilities.Ability{
	//	ID: abilityID,
	//	OwnerId: ownerID,
	//}
	//
	//abilityRequest := abilities.GetAbilityRequest{
	//	Ability: ability,
	//}
	//
	//a, err := getAbility(nil, abilityRequest)
	//
	//spew.Dump(a, err)
}

type Env struct {
	dbHost   string
	dbPort   string
	dbUser   string
	dbName   string
	httpPort string
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		//log.("Error loading .env file")
	}

	return &Env{
		dbHost:   os.Getenv("DB_HOST"),
		dbPort:   os.Getenv("DB_PORT"),
		dbUser:   os.Getenv("DB_USER"),
		dbName:   os.Getenv("DB_NAME"),
		httpPort: os.Getenv("HTTP_PORT"),
	}
}