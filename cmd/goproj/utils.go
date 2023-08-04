package main

import (
	log "github.com/sirupsen/logrus"
)

func TestResult(msg string, err error) bool {
	fmtmsg := "%-80s %s"
	if err != nil {
		log.Errorf(fmtmsg, msg, "success")
		return false
	}
	log.Debugf(fmtmsg, msg, "success")
	return true
}
