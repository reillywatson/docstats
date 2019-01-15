docstats
--------

docstats is a program to generate statistics about documentation coverage in a Go repository.

Sample usage:
```
Reillys-Mac:src reilly$ go install github.com/reillywatson/docstats/cmd/docstats
Reillys-Mac:src reilly$ docstats github.com/reillywatson/goloose
Packages: 1
Packages with docstrings: 0
Funcs: 30
Funcs with docstrings: 0
Types: 1
Types with docstrings: 0
```

Sample usage on a less-embarrassing repository:
```
Reillys-Mac:src reilly$ docstats github.com/globalsign/mgo
Packages: 14
Packages with docstrings: 5
Funcs: 917
Funcs with docstrings: 262
Types: 158
Types with docstrings: 89
```
