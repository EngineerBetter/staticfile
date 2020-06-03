package staticfile_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnit(t *testing.T) {
	suite := spec.New("staticfile", spec.Report(report.Terminal{}), spec.Parallel())
	suite("Build", testBuild)
	suite("BuildpackYAMLParser", testBuildpackYAMLParser)
	suite("Clock", testClock)
	suite("ConfigInstaller", testConfigInstaller)
	suite("Detect", testDetect)
	suite("LogEmitter", testLogEmitter)
	suite("ProfileDWriter", testProfileDWriter)
	suite.Run(t)
}
