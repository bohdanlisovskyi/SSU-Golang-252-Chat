package chatlog

import "github.com/Sirupsen/logrus"

var log = logrus.New()

type Fields map[string]interface{}

func init() {

	log.Formatter = &logrus.JSONFormatter{}
}

func Log(fields map[string]interface{}, log_type string, message ...interface{}) {
	switch log_type {
	case "error":
		log.WithFields(fields).Error(message)
	case "warning":
		log.WithFields(fields).Warning(message)
	case "fatal":
		log.WithFields(fields).Fatal(message)
	default:
		log.WithFields(fields).Info(message)
	}
}

func HandleError(err error, error_message string){

	if err != nil {

		Log(Fields{}, "error", error_message, err.Error())
	}
}