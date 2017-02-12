package config

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"path/filepath"
)

type EnvironmentConfig struct {
	Name    	string
	Port    	int
	Host    	string
	Debug		bool
	CertFile       	string
	KeyFile        	string
	Database 	DatabaseConfig
	Messaging 	string
	Redis 		RedisConfig
	Security 	SecurityConfig
	Logging 	LoggingConfig
	FCM		FCMConfig
	Minio		MinioConfig
}

type MinioConfig struct {
	EndPoint	string
	AccessKeyID	string
	SecretAccessKey	string
	UseSSL		bool
}

type FCMConfig struct {
	APIKey		string
	EndPoint	string
}

type SecurityConfig struct {
	JWT JWTConfig
	CSRF CSRFConfig
}

type JWTConfig struct {
	SigningKey string
	TokenValidityInSeconds int64
	TokenValidityInSecondsForRememberMe int64
}

type  CSRFConfig struct {
	SecretBase string
}

type DatabaseConfig struct {
	DriverName 	string
	Host    string
	Port    int
	Name 	string
	Username string	// Username login db
	Password string  //Login password of the database.
	Debug  bool
	LogFilePath  string
}

type RedisConfig struct {
	Host    string
	Port    int
}


type LoggingConfig struct {
	LogToConsole bool
	LogToFile    bool
	LogFilePath  string
	LogLevel     string
}

var Env EnvironmentConfig

func InitConf(env string) (err error) {
	//pwd, _ := os.Getwd()
	envFilename := "application.json"
	if len(env)  >  0{
		envFilename = "application-" + env + ".json"
	}

	fmt.Println("Env config:", envFilename)
	//envFilePath := path.Join(pwd, "config", envFilename)

	envFilePath := FindConfigFile(envFilename)
	contents, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, &Env)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig will try to search around for the corresponding config file.
// It will search /tmp/fileName then attempt ./config/fileName,
// then ../config/fileName and last it will look at fileName
func FindConfigFile(fileName string) string {
	if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	}

	return fileName
}
