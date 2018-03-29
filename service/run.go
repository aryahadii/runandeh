package service

import (
	"encoding/json"

	"github.com/aryahadii/runandeh/runner"
	"github.com/sirupsen/logrus"
)

// Run gets raw run request and runs
func Run(request []byte) {
	runInfo := &runner.RunRequest{}
	json.Unmarshal(request, runInfo)
	logrus.Infof("run request -> %v", runInfo)
}
