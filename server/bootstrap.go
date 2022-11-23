package server

// A centralized place where all dependencies should be declared as independent function to make
// codebase more loosely coupled.
import (
	// "acquia/decision-service/pkg/database"
	// "acquia/decision-service/pkg/rest"

	"os"

	"github.com/Scrummyy/scrummyy-api/internal/constants"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitLog initializes logrus logger instance.
func InitLog(conf *viper.Viper) {
	logrus.SetReportCaller(false)

	// set formatter
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
	})

	l := conf.GetString(constants.DebugLevel)
	lv, err := logrus.ParseLevel(l)
	if err != nil {
		panic(err)
	}

	logrus.SetLevel(lv)

	// add hooks for other logging like sumologic,statsd etc
	// logrus.AddHook()
}

// InitConfig initializes viper configuration instance.
func InitConfig() (*viper.Viper, error) {
	configPath, ok := os.LookupEnv(constants.EnvConfigPath)

	// create new config instance
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	logrus.Info("Searching for config file...")

	if !ok {
		// fallback to default location if the config is not found in the directory mentioned
		configPath = constants.DefaultConfigPath
	}

	config.AddConfigPath(configPath)
	config.WatchConfig()
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// InitDatabase initializes database connection and its dependencies.
// func InitDatabase(config *viper.Viper, name string) (database.Database, error) {
// 	var db database.Database
// 	var err error
// 	switch name {
// 	case constants.DecisionDatabase:
// 		db, err = database.NewMysqlDB(config.GetString(constants.DatabaseDecisionUri), database.DatabaseOptions{})
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// add a sql logger hook based on the config value
// 	db.AddHook(func(args ...interface{}) {
// 		if config.GetBool(constants.DebugSQLQueries) {
// 			q := args[0] // the first values is the query
// 			p := args[1] // the second value is the arguments passed to the DB function
// 			logrus.WithField("Arguments", p).WithField("Query", q).Debug("SQL Logger")
// 		}
// 	})

// 	return db, nil
// }

// InitHttpClient creates an http client using golang client.
// func InitHttpClient(config *viper.Viper) rest.HttpClientInterface {
// 	return &http.Client{
// 		Transport: &logger.LoggerTransport{
// 			Transport: http.Transport{
// 				Proxy: http.ProxyFromEnvironment,
// 				DialContext: (&net.Dialer{
// 					Timeout:   time.Duration(config.GetInt64(constants.HttpClientTimeout)) * time.Second,
// 					KeepAlive: time.Duration(config.GetInt64(constants.HttpClientTimeout)) * time.Second,
// 				}).DialContext,
// 				MaxIdleConns:          config.GetInt(constants.HttpClientMaxIdleConnections),
// 				IdleConnTimeout:       60 * time.Second,
// 				TLSHandshakeTimeout:   10 * time.Second,
// 				ExpectContinueTimeout: 1 * time.Second,
// 				ForceAttemptHTTP2:     true,
// 				MaxIdleConnsPerHost:   config.GetInt(constants.HttpClientMaxIdleConnsPerHost),
// 			},
// 		},
// 		Timeout: time.Duration(config.GetInt64(constants.HttpClientTimeout)) * time.Second,
// 	}
// }

// InitCaching initializes and instance of redis.
// func InitCaching(config *viper.Viper) (cache.Storer, error) {
// 	switch config.GetString(constants.CacheType) {
// 	case "redis":
// 		opt := cache.RedisOptions{
// 			Address:     config.GetString(constants.CacheRedisHost),
// 			Username:    config.GetString(constants.CacheRedisUsername),
// 			Password:    config.GetString(constants.CacheRedisPassword),
// 			DB:          config.GetInt(constants.CacheRedisDB),
// 			MaxRetries:  config.GetInt(constants.CacheRedisMaxRetries),
// 			Namespace:   config.GetString(constants.CacheRedisNamespace),
// 			DialTimeout: time.Duration(config.GetInt(constants.CacheRedisDialTimeout)) * time.Second,
// 			ReadTimeout: time.Duration(config.GetInt(constants.CacheRedisReadTimeout)) * time.Second,
// 			DefaultTTL:  time.Duration(config.GetInt(constants.CacheDefaultTTL)) * time.Minute,
// 		}
// 		return cache.NewRedisCache(opt)
// 	case "mysql":
// 		opt := cache.MySQLOptions{
// 			Uri:        config.GetString(constants.CacheMySQLUri),
// 			DefaultTTL: time.Duration(config.GetInt(constants.CacheDefaultTTL)) * time.Minute,
// 		}
// 		return cache.NewMySQLCache(opt)
// 	case "stub":
// 		return &cache.Stub{}, nil

// 	}

// 	return nil, errors.New(constants.ErrorInvalidCacheType)
// }
