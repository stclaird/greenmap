package config

import "os"

type MongodbConfig struct {
	Db                string
	User              string
	Password          string
	Connection_String string
}

func getMongoConfig() MongodbConfig {
	//retrun mongo configuration
	mongoconf := MongodbConfig{}
	mongoconf.Connection_String = os.Getenv("MONGO_CONNECTION_STRING")
	mongoconf.User = os.Getenv("MONGO_USER")
	mongoconf.Password = os.Getenv("MONGO_PASSWORD")
	mongoconf.Db = os.Getenv("MONGO_DB")

	return mongoconf
}
