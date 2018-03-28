package main

import (
	"github.com/spf13/viper"
)

// ServerConfigStruct ServerConfigStruct Struct
type ServerConfigStruct struct {
	Port      string
	Mode      string
	UploadDir string
	UploadURL string
}

// ServerConfig Server Config
var ServerConfig ServerConfigStruct

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	ServerConfig.Port = viper.GetString("server.port")
	ServerConfig.Mode = viper.GetString("server.mode")
	ServerConfig.UploadDir = viper.GetString("server.upload_dir")
	ServerConfig.UploadURL = viper.GetString("server.upload_url")
}
