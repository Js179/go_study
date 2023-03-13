package main

import (
	"errors"
	"github.com/Js179/logf"
	"time"
)

func main() {
	logf.Info("start...")
	logf.Infof("start... %v", time.Now().Format(time.DateTime))
	logf.Error(errors.New("END"))
}
