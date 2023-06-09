package database

import (
	"database/sql"
	"sandbox-api/config"
	"sandbox-api/logs"

	"fmt"

	_ "github.com/lib/pq"
)

const dbStarted = "Sucessfully Connected..."

type dbAccess struct {
	db *sql.DB
}

type RepoStore struct {
	Users  IUserRepo
	Items  IItemRepo
	Makers IMakerRepo
}

func NewDBAccess(cfg config.PostgresConfig) *RepoStore {
	db := connectDB(cfg, logs.NewAppLogger())
	return &RepoStore{
		Users:  newUserStore(db),
		Items:  newItemStore(db),
		Makers: newMakerStore(db),
	}

}

// InitDB sets a connection to the database
func connectDB(cfg config.PostgresConfig, log *logs.AppLogger) *sql.DB {
	psqlInfo := connectionString(cfg)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Startf("[DB] Error : %v", err)
	}
	//force to execute to check if there's an error
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Startf(" [DB] %v", dbStarted)
	return db
}

func connectionString(cfg config.PostgresConfig) string {
	switch cfg.Password {
	case "":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Name)
	default:
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	}
}
