docstats
--------

docstats is a program to generate statistics about documentation coverage in a Go repository.

Sample usage:
```
Reillys-Mac:src reilly$ go install github.com/reillywatson/docstats/cmd/docstats
Reillys-Mac:src reilly$ docstats github.com/reillywatson/goloose
Packages: 1
Packages with docstrings: 0
Funcs: 5
Funcs with docstrings: 0
Types: 2
Types with docstrings: 0
Struct fields: 1
Struct fields with docstrings: 0
```

Sample usage on a less-embarrassing repository:
```
Reillys-Mac:src reilly$ docstats go.mongodb.org/mongo-driver/mongo
Packages: 5
Packages with docstrings: 1
Funcs: 406
Funcs with docstrings: 400
Types: 85
Types with docstrings: 85
Struct fields: 255
Struct fields with docstrings: 9
```
