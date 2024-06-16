package mysqlrepo

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

var DBConfig Config

func Setup() {

	err := viper.Unmarshal(&DBConfig)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to decode into struct, %v", err))
	}
	if len(DBConfig.Host) == 0 {
		log.Fatal("DB_HOST was not set")
	}
	if len(DBConfig.UserName) == 0 {
		log.Fatal("DB_USERNAME was not set")
	}
	if len(DBConfig.Password) == 0 {
		log.Fatal("DB_PASSWORD was not set")
	}
	if len(DBConfig.Port) == 0 {
		log.Fatal("DB_PORT was not set")
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBConfig.UserName, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Kolkata")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	DBConfig.URL = dsn
}
