package context

import (
	"fmt"
	"os"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
)

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
