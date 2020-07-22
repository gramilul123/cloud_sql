package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_cloudshellProxy/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {

	const dbIP = ""
	const dbInstanceName = ""
	const dbName = ""
	const dbUserName = ""
	const dbPassword = ""

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello SQL!  Hello?")
	fmt.Fprint(w, "\n")

	const dbOpenString = dbUserName + ":" + dbPassword + "@unix(/cloudsql/" + dbInstanceName + ")/" + dbName

	db, err := sql.Open("mysql", dbOpenString)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Fprint(w, "Failed Connection"+"  "+dbOpenString)
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, err)
		return
	} else {
		fmt.Fprint(w, "SUCCESSFUL CONNECTION"+"  "+dbOpenString)
		fmt.Fprint(w, "\n")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS exercisecloudsql101 (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(100) NOT NULL, description TEXT, PRIMARY KEY (id))")
	if err != nil {
		fmt.Fprint(w, "CREATE TABLE failed:")
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, err)
		fmt.Fprint(w, "\n")
	} else {
		fmt.Fprint(w, "SUCCESSFUL CreateTable"+"  "+dbOpenString)
		fmt.Fprint(w, "\n")
	}

}
