package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	//postgres://usuario:clave@host:puerto/dbname
	connStr := "postgresql://kfgfktvb:zvebamjjpvdceundthvr@alpha.mkdb.sh:5432/ajemxuep"

	var err error
	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Fallo en la conexion", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Fallo en el Ping del sistema", err)
	}

	log.Println("Conexion establecida")
}
