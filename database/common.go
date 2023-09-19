package database

import (
	"database/sql"
	"fmt"
	"github.com/devJGuerrero/gambitnaxbeuser/models"
	"github.com/devJGuerrero/gambitnaxbeuser/secret-manager"
	"os"
)

var DB *sql.DB
var Failure error
var SecretManager models.SecretManagerRDSJson

func ReadSecretManager() error {
	SecretManager, Failure = secret_manager.GetSecret(os.Getenv("secret-name"))
	return Failure
}

func prepareStringConnect(keys models.SecretManagerRDSJson) string {
	var dbHost, dbName, dbUser, authToken string
	dbHost = keys.Host
	dbName = "gambitnaxbe"
	dbUser = keys.Username
	authToken = keys.Password
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbHost, dbName)
	fmt.Println(dsn)
	return dsn
}

func Connect() error {
	DB, Failure = sql.Open("mysql", prepareStringConnect(SecretManager))
	if Failure != nil {
		fmt.Println("Error. Error. MySQL connection not established. " + Failure.Error())
		return Failure
	}
	Failure = DB.Ping()
	if Failure != nil {
		fmt.Println("Error. Error. MySQL connection not established ping. " + Failure.Error())
		return Failure
	}
	fmt.Println("MySQL connection successfully established.")
	return nil
}
