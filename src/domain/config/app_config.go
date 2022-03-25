package config

import "os"

type AppConfig struct {
	AppName              string
	Env                  string
	MongoCnn             string
	MongoDbName          string
	Port                 string
	Version              string
	BaseUrl              string
	SecretTokenKey       string
	DocsPath             string
	ServerUrl            string
	ImageDefault         string
	DefaultUserImageName string
}

func (c AppConfig) SetFromEnvFile() *AppConfig {
	return &AppConfig{
		Env:                  os.Getenv("ENV"),
		MongoCnn:             os.Getenv("MONGO_CNN"),
		Port:                 os.Getenv("PORT"),
		MongoDbName:          os.Getenv("MONGO_DB_NAME"),
		AppName:              os.Getenv("APP_NAME"),
		Version:              os.Getenv("VERSION"),
		BaseUrl:              os.Getenv("BASEURL"),
		SecretTokenKey:       os.Getenv("TOKEN_KEY"),
		DocsPath:             os.Getenv("DOCS_PATH"),
		ServerUrl:            os.Getenv("SERVER_URL"),
		ImageDefault:         os.Getenv("IMAGE_DEFAULT"),
		DefaultUserImageName: os.Getenv("DEFAULT_USER_IMAGE_NAME"),
	}
}
