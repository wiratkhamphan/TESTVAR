package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Env struct {
	DBdriver string
	DBuser   string
	DBpass   string
	DBname   string
}

func ConnfileEnv() (*Env, error) {
	envPath := filepath.Join("./sec/data", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal("Error Logdinh .env file :", err)
	}

	Env := &Env{
		DBdriver: os.Getenv("DB_Driver"),
		DBuser:   os.Getenv("DB_User"),
		DBpass:   os.Getenv("DB_Pass"),
		DBname:   os.Getenv("DB_Name"),
	}
	// Test connect file mysql.env
	// fmt.Println(Env)
	return Env, nil
}

func DBConnection() (*sql.DB, error) {
	env, err := ConnfileEnv()
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
	// defer db.Close()
	return db, err
}
