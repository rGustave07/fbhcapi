package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	application := App{}
	application.Initialize().Run(":8080")
}
