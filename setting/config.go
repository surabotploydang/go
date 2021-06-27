package setting

import (
	"crypto/rsa"
	"github.com/patcharp/golib/cache"
	"github.com/patcharp/golib/util"
)

var (
	Production    = util.GetEnv("SERVER_MODE", "prod") == "prod"
	AppName       string
	AppVersion    string
	ListeningPort string

	VerifyKey *rsa.PublicKey

	// Datastore
	//MophDb = database.Config{
	//	Host:      util.GetEnv("MOPHIC_DB_HOST", "127.0.0.1"),
	//	Port:      util.GetEnv("MOPHIC_DB_PORT", "3306"),
	//	Username:  util.GetEnv("MOPHIC_DB_USER", ""),
	//	Password:  util.GetEnv("MOPHIC_DB_PASSWORD", ""),
	//	Name:      util.GetEnv("MOPHIC_DB_NAME", ""),
	//	DebugMode: util.GetEnv("SERVER_MODE", "prod") != "prod",
	//}
	//QueueDb = database.Config{
	//	Host:      util.GetEnv("QUEUE_DB_HOST", "127.0.0.1"),
	//	Port:      util.GetEnv("QUEUE_DB_PORT", "3306"),
	//	Username:  util.GetEnv("QUEUE_DB_USER", ""),
	//	Password:  util.GetEnv("QUEUE_DB_PASSWORD", ""),
	//	Name:      util.GetEnv("QUEUE_DB_NAME", ""),
	//	DebugMode: util.GetEnv("SERVER_MODE", "prod") != "prod",
	//}

	// Redis Caching
	Caching = cache.Config{
		Host:     util.GetEnv("CACHING_HOST", "127.0.0.1"),
		Port:     util.GetEnv("CACHING_PORT", "6379"),
		Password: util.GetEnv("CACHING_PASSWORD", ""),
		Db:       util.AtoI(util.GetEnv("CACHING_DB", "2"), 2),
	}
)
