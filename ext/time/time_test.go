package time_test

import (
	"testing"

	arrow "github.com/bmuller/arrow/lib"

	"github.com/gopackt/jinjia2"
	"github.com/gopackt/jinjia2/ext/time"
	tu "github.com/gopackt/jinjia2/testutils"
)

func Env(root string) *jinjia2.Environment {
	env := tu.TestEnv(root)
	env.Statements.Update(time.Statements)
	cfg := time.NewConfig()
	parsed, _ := arrow.CParse("%Y-%m-%d %H:%M:%S", "1984-06-07 16:40:00")
	cfg.Now = &parsed
	env.Config.Ext["time"] = cfg
	return env
}

func TestTimeStatement(t *testing.T) {
	root := "./testData"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}
