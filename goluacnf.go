package goluacnf

import (
	"errors"

	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type Data map[interface{}]interface{}

type Config struct {
	data Data
}

func (c *Config) GetData() Data {
	return c.data
}

func (c *Config) SetData(d Data) {
	c.data = d
}

func (c *Config) String(key string) string {
	return c.GetData()[string(key)].(string)
}

func (c *Config) Int(key string) int {
	return c.GetData()[string(key)].(int)
}

func NewConfig() Config {
	return Config{}
}

func Register(filepath, environ string) (Config, error) {
	conf := NewConfig()

	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(filepath); err != nil {
		return conf, err
	}

	table := L.GetGlobal(environ).(*lua.LTable)
	mp, err := mapper(table)
	if err != nil {
		return conf, err
	}

	conf.SetData(mp)

	return conf, nil
}

func mapper(tbl *lua.LTable) (map[interface{}]interface{}, error) {
	opt := gluamapper.Option{NameFunc: gluamapper.ToUpperCamelCase}

	mp, ok := gluamapper.ToGoValue(tbl, opt).(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("arguments #1 must be a table, but got an array")
	}
	return mp, nil
}
