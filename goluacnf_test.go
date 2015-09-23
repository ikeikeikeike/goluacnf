package goluacnf

import (
	"path"
	"runtime"
	"testing"

	"github.com/kr/pretty"
)

func testcnf() (Config, error) {
	return Register(path.Join(Root, "testcnf.lua"), Env)
}

func TestInit(t *testing.T) {
	cnf, err := testcnf()
	checkFatal(t, err)

	pretty.Println(cnf)
}

// func TestEnv(t *testing.T) {
// cnf, err := Register(
// path.Join(Root, "testcnf.lua"),
// Env,
// )
// checkFatal(t, err)

// pretty.Println(conf)
// }

func checkFatal(t *testing.T, err error) {
	if err == nil {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Fatal()
	}

	t.Fatalf("Fail at %v:%v; %v", file, line, err)
}
