package calculator

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}
var Got float64

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                "calculadora",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		Got = 0 // clean the state before every scenario
	})

	ctx.Step(`^numbers (\d.\d+) and (\d.\d+)$`, numbersAnd)
	ctx.Step(`^sum should be (\d.\d+)$`, sumShouldBe)
}

func numbersAnd(arg1, arg2 float64) error {
	Got = roundToTwoDecimalPoints(Sum(arg1, arg2))
	return nil
}

func sumShouldBe(want float64) error {
	if Got != want {
		return fmt.Errorf("got %v want %v", Got, want)
	}
	return nil
}
