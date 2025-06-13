package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/aglis-lab/golem/src/app"
)

func TestMain(m *testing.M) {
	os.Chdir("../../../")

	app.Init(context.Background())

	exitVal := m.Run()

	os.Exit(exitVal)
}
