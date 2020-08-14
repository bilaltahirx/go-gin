package env

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type envFile struct {
	DbName             string
	DbUsername         string
	DbPassword         string
	DbHost             string
	DbPort             string
	BuildEnv           string
	ServerPort         string
	DbPoolSize         int
	CadenceService     string
	TaskListName       string
	CadenceDomain      string
	ClientName         string
	CadenceServiceName string
	NatsAddress        string
	NatsCluster        string
	NatsClient         string
	PromProtocol       string
	PromHost           string
	PromPort           string
	IsContainer        bool
}

func (e *envFile) GetAddr() string {
	return e.DbHost + ":" + e.DbPort
}

var Env *envFile

func init() {
	_ = godotenv.Load()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//IsContainer, _ := strconv.ParseBool(os.Getenv("IS_CONTAINER"))
	//DbPoolSize, _ := strconv.Atoi(os.Getenv("DB_POOL_SIZE"))

	logger.Info("I am here just checking things out")
	Env = &envFile{
		DbName:             os.Getenv("DB_NAME"),
		DbUsername:         os.Getenv("DB_USERNAME"),
		DbPassword:         os.Getenv("DB_PASSWORD"),
		DbHost:             os.Getenv("DB_HOST"),
		DbPort:             os.Getenv("DB_PORT"),
		//BuildEnv:           os.Getenv("BUILD_ENV"),
		//ServerPort:         os.Getenv("SERVER_PORT"),
		//CadenceService:     os.Getenv("CADENCE_SERVICE"),
		//TaskListName:       os.Getenv("TASK_LIST_NAME"),
		//CadenceDomain:      os.Getenv("CADENCE_DOMAIN"),
		//ClientName:         os.Getenv("CLIENT_NAME"),
		//CadenceServiceName: os.Getenv("CADENCE_SERVICE_NAME"),
		//NatsAddress:        os.Getenv("NATS_URL"),
		//NatsCluster:        os.Getenv("NATS_CLUSTER_ID"),
		//NatsClient:         os.Getenv("NATS_CLIENT_ID"),
		//PromProtocol:       os.Getenv("PROM_PROTOCOL"),
		//PromHost:           os.Getenv("PROM_HOST"),
		//PromPort:           os.Getenv("PROM_PORT"),
		//DbPoolSize:         DbPoolSize,
		//IsContainer:        IsContainer,
	}
}
