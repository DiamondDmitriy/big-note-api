package database

import (
	"database/sql"
	"fmt"
	"github.com/DiamondDmitriy/big-note-api/config"
	_ "github.com/lib/pq"
)

func NewDB(cnf *config.DB) (*sql.DB, error) {
	// todo: настроить sslmode
	connInfo := fmt.Sprintf(
		"host=%s dbname=%s port=%s user=%s password=%s sslmode=disable",
		cnf.Host,
		cnf.Database,
		cnf.Port,
		cnf.Username,
		cnf.Password,
	)
	return sql.Open(cnf.DriverName, connInfo)
}
