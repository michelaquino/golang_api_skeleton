package context

import (
	"fmt"
	"io"
	"os"
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
	mongoURL := "localhost"
	mongoPort := 27017
	mongoAddress := fmt.Sprintf("%s:%d", mongoURL, mongoPort)
	mongoTimeout := time.Duration(60) * time.Second

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs: []string{
			mongoAddress,
		},
		Database: "api",
		Username: "",
		Password: "",
		Timeout:  mongoTimeout,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	return session, err
}

func getNewLogInstance() *logrus.Logger {
	logrusLog := logrus.New()
	logrusLog.Level = getLogLevel()
	logrusLog.Out = os.Stdout
	return logrusLog
}

func getLogLevel() logrus.Level {
	logLevelConfig := "debug"
	level, err := logrus.ParseLevel(logLevelConfig)
	if err != nil {
		return logrus.InfoLevel
	}

	return level
}

func getLogFile() io.Writer {
	logFileName := "api.log"
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		errorMsg := fmt.Sprintf("Error on open log file: %s", err.Error())
		panic(errorMsg)
	}

	return logFile
}
