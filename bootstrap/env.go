package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	Host                   string `mapstructure:"DB_HOST"`
	Port                   string `mapstructure:"DB_PORT"`
	User                   string `mapstructure:"DB_USER"`
	Password               string `mapstructure:"DB_PASS"`
	Database               string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.AutomaticEnv() // Otomatis membaca variabel lingkungan

	// Membaca nilai variabel lingkungan menggunakan Viper
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")

	// Menampilkan nilai variabel lingkungan
	fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_PORT: %s\n", dbPort)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	fmt.Println(env)

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
