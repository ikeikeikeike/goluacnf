package goluacnf

import (
	"errors"

	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type Data map[interface{}]interface{}

type Config struct {
	data Data
	tbl  *lua.LTable
}

func (c *Config) GetData() Data {
	return c.data
}

func (c *Config) Get(key string) interface{} {
	return c.GetData()[string(key)]
}

func (c *Config) String(key string) string {
	return c.Get(key).(string)
}

func (c *Config) Int(key string) int {
	return int(c.Get(key).(float64))
}

func (c *Config) Int64(key string) int64 {
	return int64(c.Int(key))
}

func (c *Config) Float(key string) float64 {
	return c.Get(key).(float64)
}

func (c *Config) Float32(key string) float32 {
	return float32(c.Float(key))
}

func (c *Config) Bool(key string) bool {
	return c.Get(key).(bool)
}

func (c *Config) Map(st interface{}) error {
	return gluamapper.Map(c.tbl, st)
}

func NewConfig(table *lua.LTable) (Config, error) {
	mp, err := mapper(table)
	if err != nil {
		return Config{}, err
	}

	return Config{data: mp, tbl: table}, nil
}

func Register(filepath, environ string) (Config, error) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(filepath); err != nil {
		return Config{}, err
	}

	return NewConfig(L.GetGlobal(environ).(*lua.LTable))
}

func mapper(tbl *lua.LTable) (map[interface{}]interface{}, error) {
	opt := gluamapper.Option{NameFunc: gluamapper.ToUpperCamelCase}

	mp, ok := gluamapper.ToGoValue(tbl, opt).(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("arguments #1 must be a table, but got an array")
	}
	return mp, nil
}
