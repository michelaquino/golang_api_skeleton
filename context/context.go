package context

import (
	"fmt"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

type context struct {
	mongoSession *mgo.Session
	log          Logger
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

func (obj *context) GetLogger() Logger {
	return obj.log
}

func newContextInstance() *context {
	session, err := getNewMongoSession()
	if err != nil {
		errorMsg := fmt.Sprintf("Error on start database: %s", err.Error())
		panic(errorMsg)
	}

	logrusLog := GetLogger()
	return &context{
		mongoSession: session,
		log:          logrusLog,
	}
}
