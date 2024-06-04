package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Env struct {
	DBdriver string
	DBuser   string
	DBpass   string
	DBname   string
}

func ConnfileEnv_book() (*Env, error) {
	envPath := filepath.Join("./sec/data/data_shop_book", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal("Error Logdinh .env file :", err)
	}

	Env := &Env{
		DBdriver: os.Getenv("DB_Driver"),
		DBuser:   os.Getenv("DB_User"),
		DBpass:   os.Getenv("DB_Pass"),
		DBname:   os.Getenv("DB_Name1"),
	}
	// Test connect file mysql.env
	// fmt.Println(Env)
	return Env, nil
}

func DBConnection_book() (*sql.DB, error) {
	env, err := ConnfileEnv_book()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(env.DBdriver, fmt.Sprintf("%s:%s@/%s", env.DBuser, env.DBpass, env.DBname))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("ConnfileEnv_book", env)
	// defer db.Close()
	return db, err

}
