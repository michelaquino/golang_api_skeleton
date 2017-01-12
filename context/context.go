package context

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type context struct {
	mongoSession *mgo.Session
	log          *logrus.Logger
}

var contextInstance *context
var once sync.Once

// GetAPIContext return a instance of the api context
func GetAPIContext() *context {
	once.Do(func() {
		contextInstance = newContextInstance()
	})

	return contextInstance
}

func (obj *context) GetMongoSession() *mgo.Session {
	return obj.mongoSession.Copy()
}

func (obj *context) GetLogger() *logrus.Logger {
	return obj.log
}

func newContextInstance() *context {
	session, err := getNewMongoSession()
	if err != nil {
		errorMsg := fmt.Sprintf("Error on start database: %s", err.Error())
		panic(errorMsg)
	}

	logrusLog := getNewLogInstance()

	return &context{
		mongoSession: session,
		log:          logrusLog,
	}
}

func getNewMongoSession() (*mgo.Session, error) {
	mongoURL := os.Getenv("MONGO_URL")
	mongoPort := getMongoPort()
	mongoTimeout := getMongoTimeout()
	mongoAddress := fmt.Sprintf("%s:%d", mongoURL, mongoPort)

	mongoDatabaseName := os.Getenv("MONGO_DATABASE_NAME")
	mongoUserName := os.Getenv("MONGO_DATABASE_USERNAME")
	mongoPassword := os.Getenv("MONGO_DATABASE_PASSWORD")

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs: []string{
			mongoAddress,
		},
		Database: mongoDatabaseName,
		Username: mongoUserName,
		Password: mongoPassword,
		Timeout:  mongoTimeout,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	return session, err
}

func getMongoPort() int {
	mongoPort, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		return 27017
	}

	return mongoPort
}

func getMongoTimeout() time.Duration {
	mongoTimeout, err := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	if err != nil {
		return time.Duration(60) * time.Second
	}

	return time.Duration(mongoTimeout) * time.Second
}

func getNewLogInstance() *logrus.Logger {
	logrusLog := logrus.New()
	logrusLog.Level = getLogLevel()
	logrusLog.Out = getLogOut()
	return logrusLog
}

func getLogLevel() logrus.Level {
	logLevelConfig := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(logLevelConfig)
	if err != nil {
		return logrus.DebugLevel
	}

	return level
}

func getLogOut() io.Writer {
	sendLogToStdout := false
	if logToStdout, err := strconv.ParseBool(os.Getenv("LOG_TO_STDOUT")); err == nil {
		sendLogToStdout = logToStdout
	}

	logFileName := os.Getenv("LOG_FILE_NAME")
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error on open log file: %s", err.Error())
		return os.Stdout
	}

	if sendLogToStdout {
		return io.MultiWriter(os.Stdout, logFile)
	}

	return logFile
}
