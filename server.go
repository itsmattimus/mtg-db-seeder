package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	m "simianspiritguide/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	oracle m.OracleCards
	cc     int
)

func main() {
	sqlDB, err := sql.Open("mysql", "root:mysql1@tcp(127.0.0.1)/scryfall-bulk-data")
	checkError(err)

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "scryfall-bulk-data.",
			SingularTable: false,
		},
	})
	checkError(err)

	defer sqlDB.Close()
	fmt.Println("connected!")

	plan, _ := os.Open("bulk-card-json/oracle-cards.json")
	dec := json.NewDecoder(io.Reader(plan))

	gdb.AutoMigrate(&oracle)

	t, ft := dec.Token()
	checkError(ft)
	fmt.Printf("%T: %v\n", t, t)

	for dec.More() {
		cc++
		err := dec.Decode(&oracle)
		checkError(err)

		create := gdb.Create(&oracle)
		checkError(create.Error)
	}

	t, lt := dec.Token()
	checkError(lt)
	fmt.Printf("%T: %v\n", t, t)

	fmt.Printf("parsed %d cards!\n", cc)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
