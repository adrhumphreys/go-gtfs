package internal

import (
	"fmt"
	csvtag "github.com/artonge/go-csv-tag"
	"github.com/jmoiron/sqlx"
	"gtfs-cli/internal/dbx"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
	"sync"
)

type Importer struct {
	model string
	table string
	db    *sqlx.DB
}

func loadInMemory(folderPath string) GTFS {
	var GTFS = GTFS{}

	filesToLoad := map[string]interface{}{
		"agency.txt":   &GTFS.agencies,
		"routes.txt":   &GTFS.routes,
		"calendar.txt": &GTFS.calendars,
		"calendar_dates.txt": &GTFS.calendarDates,
		"shapes.txt":         &GTFS.shapes,
	}

	for file, dest := range filesToLoad {
		filePath := path.Join(folderPath, file)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			fmt.Printf("Missing file: %v\n", filePath)
			continue
		}

		err = csvtag.LoadFromPath(filePath, dest)
		if err != nil {
			log.Fatalf("Error loading file (%v)\n	==> %v", file, err)
		}
	}

	return GTFS
}

func Import(folderPath string) {
	GTFS := loadInMemory(folderPath)
	db := dbx.Connect()
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		count := 0
		for _, agency := range GTFS.agencies {
			importModel(agency, "agency", db)
			count++
		}
		fmt.Printf("🛵 %v agencies imported\n", count)

		count = 0
		for _, route := range GTFS.routes {
			importModel(route, "route", db)
			count++
		}
		fmt.Printf("🛵 %v routes imported\n", count)
		wg.Done()
	}(&wg)


	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		count := 0
		for _, calendar := range GTFS.calendars {
			importModel(calendar, "calendar", db)
			count++
		}
		fmt.Printf("🛵 %v calendars imported\n", count)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		count := 0
		for _, calendarDate := range GTFS.calendarDates {
			importModel(calendarDate, "calendar_date", db)
			count++
		}
		fmt.Printf("🛵 %v calendar dates imported\n", count)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		count := 0
		for _, shape := range GTFS.shapes {
			importModel(shape, "shape", db)
			count++
		}
		fmt.Printf("🛵 %v shapes imported\n", count)
		wg.Done()
	}(&wg)

	wg.Wait()

}

func importModel(model interface{}, tableName string, db *sqlx.DB) {
	fields := getDBFields(model)
	columns := getInsert(fields, false)
	fieldTagBinding := getInsert(fields, true)
	sql := "INSERT INTO " + tableName + " (" + columns + ") VALUES (" + fieldTagBinding + ")"
	_, err := sqlx.NamedExec(db, sql, model)
	if err != nil {
		log.Fatal(err)
	}
}

func getDBFields(model interface{}) []string {
	var dbFields []string
	lookup := reflect.TypeOf(model)

	for i := 0; i < lookup.NumField(); i++ {
		tag := lookup.Field(i).Tag.Get("db")
		dbFields = append(dbFields, tag)
	}

	return dbFields
}

func getInsert(values []string, prependColon bool) string {
	var columnFields string

	var prependString = ""
	if prependColon {
		prependString = ":"
	}

	for _, field := range values {
		columnFields += prependString + field + ","
	}

	columnFields = strings.TrimRight(columnFields, ",")

	return columnFields
}
