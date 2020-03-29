package main

import log "github.com/sirupsen/logrus"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	log.WithFields(log.Fields{
		"plugin": "kafka",
		"size":   10,
	}).Info("Hello World!")
}
