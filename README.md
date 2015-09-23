go-luacnf
========

[![Build Status](https://travis-ci.org/ikeikeikeike/goluacnf.png)](https://travis-ci.org/ikeikeikeike/goluacnf)

> ## wats go-luacnf
> Just I wanted to do that.

```Go
import "github.com/ikeikeikeike/goluacnf"
```

## Usage 

We can use lua file as a configuration file like this.

```lua
function merge(t1, t2)
    local t = {}
    for k, v in pairs(t1) do t[k] = v end
    for k, v in pairs(t2) do t[k] = v end
    return t
end

common = {
    letter_a = "a", -- site title
    letter_b = "b",
    letter_c = "c",
    string = "1",
    int = 1,
    float = 0.01,
    yes = true,
    no = false,
    map = {{
        first = "common1",
        second = "common1",
        third = "common1",
    },{
        first = "common2",
        second = "common2",
        third = "common2",
    }}
}

development = merge(common, {
    name = "development",
    dsn = "postgres://username:password@development.local:5432/goluacnf",
    table = {{
        first = "development1",
        second = "development1",
        third = "development1",
    },{
        first = "development2",
        second = "development2",
        third = "development2",
    }}
})

production = merge(common, {
    name = "production",
    dsn = "postgres://username:password@production.local:5432/goluacnf",
    table = {{
        first = "production1",
        second = "production1",
        third = "production1",
    },{
        first = "production2",
        second = "production2",
        third = "production2",
    }}
})
```

### Loading a configuration file

We can write in this way if we want to change configuration's file path.

```go
func (app *App) initConfig() {
	cnf, err := goluacnf.Register(
		path.Join(goluacnf.Root, "configs/config.lua"),
		goluacnf.Env,
	)
	if err != nil {
		panic(err)
	}

	app.Conf = cnf
}
```

### Changing a environment


