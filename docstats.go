package docstats

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func StatsForDir(dir string) (PkgStats, error) {
	var stats PkgStats
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if info != nil && info.IsDir() {
			stats = stats.Add(parseDir(path))
		}
		return nil
	})
	return stats, err
}

type PkgStats struct {
	Pkgs                int
	PkgsWithDoc         int
	Funcs               int
	FuncsWithDoc        int
	Types               int
	TypesWithDoc        int
	StructFields        int
	StructFieldsWithDoc int
}

func (s PkgStats) Add(o PkgStats) PkgStats {
	return PkgStats{
		Pkgs:                s.Pkgs + o.Pkgs,
		PkgsWithDoc:         s.PkgsWithDoc + o.PkgsWithDoc,
		Funcs:               s.Funcs + o.Funcs,
		FuncsWithDoc:        s.FuncsWithDoc + o.FuncsWithDoc,
		Types:               s.Types + o.Types,
		TypesWithDoc:        s.TypesWithDoc + o.TypesWithDoc,
		StructFields:        s.StructFields + o.StructFields,
		StructFieldsWithDoc: s.StructFieldsWithDoc + o.StructFieldsWithDoc,
	}
}

func (s PkgStats) String() string {
	return fmt.Sprintf(`Packages: %d
Packages with docstrings: %d (%.2f%%)
Funcs: %d
Funcs with docstrings: %d (%.2f%%)
Types: %d
Types with docstrings: %d (%.2f%%)
Struct fields: %d
Struct fields with docstrings: %d (%.2f%%)
`, s.Pkgs, s.PkgsWithDoc, percent(s.PkgsWithDoc, s.Pkgs), s.Funcs, s.FuncsWithDoc, percent(s.FuncsWithDoc, s.Funcs), s.Types, s.TypesWithDoc, percent(s.TypesWithDoc, s.Types), s.StructFields, s.StructFieldsWithDoc, percent(s.StructFieldsWithDoc, s.StructFields))
}

func percent(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b) * 100
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
			for _, spec := range t.Decl.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					switch innerType := spec.Type.(type) {
					case *ast.StructType:
						if innerType.Fields != nil {
							for _, field := range innerType.Fields.List {
								stats.StructFields++
								if field.Doc != nil {
									stats.StructFieldsWithDoc++
								}
							}
						}
					}
				}
			}
		}
	}
	return stats
}
