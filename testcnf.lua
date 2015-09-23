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
    },{
        first = "development3",
        second = "development3",
        third = "development3",
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
    },{
        first = "production3",
        second = "production3",
        third = "production3",
    }}

})

overwrite = merge(common, {
    name = "overwrite",
    dsn = "postgres://username:password@overwrite.local:5432/goluacnf",
    table = {{
        first = "overwrite1",
        second = "overwrite1",
        third = "overwrite1",
    },{
        first = "overwrite2",
        second = "overwrite2",
        third = "overwrite2",
    },{
        first = "overwrite3",
        second = "overwrite3",
        third = "overwrite3",
    }},
    letter_a = "overwrited letter a", -- site title
    letter_b = "overwrited letter b",
    letter_c = "overwrited letter c",
})
