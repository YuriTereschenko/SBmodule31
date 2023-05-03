package app

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"main/internal/controller"
	"main/internal/repository/postgres"
	"main/internal/usecase"
)

func Run() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=goLANGdb")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	srv := postgres.Storage{DB: db}
	DBInter := usecase.NewDBinter(&srv)
	var port string
	flag.StringVar(&port, "port", "", "port to connect")
	flag.Parse()
	controller.RunHttp("localhost:"+port, DBInter)
}
