package types_test

import (
	"os"
	"testing"

	"github.com/four4two/fury/app"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}
