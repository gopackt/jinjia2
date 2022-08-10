package django_test

import (
	"testing"

	"github.com/gopackt/jinjia2"
	"github.com/gopackt/jinjia2/ext/django"
	tu "github.com/gopackt/jinjia2/testutils"
)

func Env(root string) *jinjia2.Environment {
	env := tu.TestEnv(root)
	env.Filters.Update(django.Filters)
	env.Statements.Update(django.Statements)
	return env
}

func TestDjangoTemplates(t *testing.T) {
	root := "./testData"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}

func TestDjangoFilters(t *testing.T) {
	root := "./testData/filters"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}

func TestDjangoStatements(t *testing.T) {
	root := "./testData/statements"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}
