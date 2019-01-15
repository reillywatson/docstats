package docstats

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func StatsForDir(dir string) (PkgStats, error) {
	var stats PkgStats
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			stats = stats.Add(parseDir(path))
		}
		return nil
	})
	return stats, err
}

type PkgStats struct {
	Pkgs         int
	PkgsWithDoc  int
	Funcs        int
	FuncsWithDoc int
	Types        int
	TypesWithDoc int
}

func (s PkgStats) Add(o PkgStats) PkgStats {
	return PkgStats{
		Pkgs:         s.Pkgs + o.Pkgs,
		PkgsWithDoc:  s.PkgsWithDoc + o.PkgsWithDoc,
		Funcs:        s.Funcs + o.Funcs,
		FuncsWithDoc: s.FuncsWithDoc + o.FuncsWithDoc,
		Types:        s.Types + o.Types,
		TypesWithDoc: s.TypesWithDoc + o.TypesWithDoc,
	}
}

func (s PkgStats) String() string {
	return fmt.Sprintf(`Packages: %d
Packages with docstrings: %d
Funcs: %d
Funcs with docstrings: %d
Types: %d
Types with docstrings: %d
`, s.Pkgs, s.PkgsWithDoc, s.Funcs, s.FuncsWithDoc, s.Types, s.TypesWithDoc)
}

func parseDir(path string) PkgStats {
	stats := PkgStats{}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, p := range pkgs {
		pkg := doc.New(p, path, 0)
		stats.Pkgs++
		if pkg.Doc != "" {
			stats.PkgsWithDoc++
		}
		for _, fn := range pkg.Funcs {
			stats.Funcs++
			if fn.Doc != "" {
				stats.FuncsWithDoc++
			}
		}
		for _, t := range pkg.Types {
			stats.Types++
			if t.Doc != "" {
				stats.TypesWithDoc++
			}
			for _, fn := range t.Funcs {
				stats.Funcs++
				if fn.Doc != "" {
					stats.FuncsWithDoc++
				}
			}
			for _, m := range t.Methods {
				stats.Funcs++
				if m.Doc != "" {
					stats.FuncsWithDoc++
				}
			}
		}
	}
	return stats
}
