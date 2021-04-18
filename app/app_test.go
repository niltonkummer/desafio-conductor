package app

import (
	"fmt"
	"os"
	"testing"

	"github.com/niltonkummer/desafio-conductor/app/config"

	"gorm.io/driver/sqlite"
)

var a *app

func TestMain(t *testing.M) {
	os.Exit(testMain(t))
}

func testMain(t *testing.M) int {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	a = NewApplication(&config.Config{
		DBDialect: sqlite.Open("../db/test_data.sqlite"),
	})

	return t.Run()
}

func seedTestDB() {
	// TODO seed
}
