package constant

import "github.com/patcharp/golib/util"

var (
	SERVER_MODE = util.GetEnv("SERVER_MODE", "")
	// // Caching
	CACHING_HOST     = util.GetEnv("CACHING_HOST", "")
	CACHING_PORT     = util.GetEnv("CACHING_PORT", "")
	CACHING_PASSWORD = util.GetEnv("CACHING_PASSWORD", "")
	// // ONEID
	ONEID_HOST = util.GetEnv("ONEID_HOST", "127.0.0.1")
	// // SIGN
	SIGN_HOST         = util.GetEnv("SIGN_HOST", "")
	SIGN_CREDENTIALID = util.GetEnv("SIGN_CREDENTIALID", "")
)
