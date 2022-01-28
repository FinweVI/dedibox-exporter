package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestSetLogLevel(t *testing.T) {
	for i := 0; i <= 4; i++ {
		setLogLevel(i)
		switch i {
		case 0:
			if !log.IsLevelEnabled(log.DebugLevel) {
				t.Fatal("DebugLevel not enabled")
			}
		case 1:
			if !log.IsLevelEnabled(log.InfoLevel) {
				t.Fatal("InfoLevel not enabled")
			}
		case 2:
			if !log.IsLevelEnabled(log.WarnLevel) {
				t.Fatal("WarnLevel not enabled")
			}
		case 3:
			if !log.IsLevelEnabled(log.ErrorLevel) {
				t.Fatal("ErrorLevel not enabled")
			}
		case 4:
			if !log.IsLevelEnabled(log.FatalLevel) {
				t.Fatal("FatalLevel not enabled")
			}
		}
	}
}
