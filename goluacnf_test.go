package goluacnf

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func getConfig() (Config, error) {
	return Register(path.Join(Root, "testcnf.lua"), Env)
}

func TestInit(t *testing.T) {
	_, err := getConfig()
	checkFatal(t, err)
}

func TestEnv(t *testing.T) {
	if Env != "development" {
		t.Fatalf("Fatal default environment: Didn't change environment to `development` thats keeping `%v` now.", Env)
	}

	os.Setenv("GOLUACNF_ENV", "production")
	setENV(os.Getenv("GOLUACNF_ENV"))

	if Env != "production" {
		t.Errorf("Fatal changed environment: Didn't change environment to `production` thats still keeping `%v`.", Env)
	}

	os.Setenv("GOLUACNF_ENV", "development")
	setENV(os.Getenv("GOLUACNF_ENV"))
}

type testcnf struct {
	Name  string
	Dsn   string
	Table []struct {
		First  string
		Second string
		Third  string
	}

	LetterA string
	LetterB string
	LetterC string
	String  string
	Int     int
	Float   float64
	Yes     bool
	No      bool
	Map     []struct {
		First  string
		Second string
		Third  string
	}
}

func TestTypes(t *testing.T) {
	cnf, _ := getConfig()

	if "1" != cnf.String("String") {
		t.Fatalf("Fatal String type: Didn't cast to `string` thats `%v`.", cnf.String("String"))
	}
	if int(1) != cnf.Int("Int") {
		t.Fatalf("Fatal String type: Didn't cast to `int` thats `%v`.", cnf.Int("Int"))
	}
	if int64(1) != cnf.Int64("Int") {
		t.Fatalf("Fatal String type: Didn't cast to `int64` thats `%v`.", cnf.Int("Int"))
	}
	if float64(0.01) != cnf.Float("Float") {
		t.Fatalf("Fatal String type: Didn't cast to `float` thats `%v`.", cnf.Float("Float"))
	}
	if float32(0.01) != cnf.Float32("Float") {
		t.Fatalf("Fatal String type: Didn't cast to `float32` thats `%v`.", cnf.Float32("Float"))
	}
	if true != cnf.Bool("Yes") {
		t.Fatalf("Fatal String type: Didn't cast to `bool(true)` thats `%v`.", cnf.Bool("Yes"))
	}
	if false != cnf.Bool("No") {
		t.Fatalf("Fatal String type: Didn't cast to `bool(false)` thats `%v`.", cnf.Bool("No"))
	}

	var c = new(testcnf)
	cnf.Map(&c)
	if 1 != c.Int {
		t.Fatalf("Fatal Struct type: Didn't cast to `struct` thats `%v`.", c)
	}
}

func TestPool(t *testing.T) {
	Pool = Config{}
	_, _ = getConfig()

	if _, ok := Pool.GetData()["Name"]; !ok {
		t.Fatalf("Fatal Data pooling: `%v`.", Pool)
	}
}

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
