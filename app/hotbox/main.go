package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/luthfiswees/hotbox/db"
	"github.com/luthfiswees/hotbox/handler"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	
	db.CreateDB()
	db.MigrateDB()

	http.HandleFunc("/store", handler.StoreHandler)
	http.HandleFunc("/get-report", handler.GetHandler)

	fmt.Println("Now serving on " + os.Getenv("HOTBOX_PORT"))
	http.ListenAndServe(":"+os.Getenv("HOTBOX_PORT"), nil)
}
