package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	PkgName      = "PkgName"
	Type         = "Type"
	IteratorName = "IteratorName"
	Dir          = "dir"
	defaultDir   = "."
)

func main() {
	data, err := fetchDataFromArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = generateFile(data)
	fmt.Println(err)
}

func fetchDataFromArgs(args []string) (map[string]string, error) {
	data := make(map[string]string)

	for validateArgs(args) {
		if err := setData(args[:2], data); err != nil {
			return nil, err
		}

		args = args[2:]
	}

	_, ok := data[PkgName]
	if !ok {
		return nil, fmt.Errorf("package name is not set")
	}

	iterType, ok := data[Type]
	if !ok {
		return nil, fmt.Errorf("type is not set")
	}

	data[IteratorName] = fmt.Sprint(strings.ToUpper(iterType[:1]), iterType[1:], "Iterator")
	return data, nil
}

func validateArgs(args []string) bool {
	size := len(args)

	if size == 0 {
		return false
	}
	if size%2 != 0 {
		return false
	}

	return true
}

func setData(args []string, data map[string]string) error {
	switch args[0] {
	case `-t`, `-type`:
		data[Type] = args[1]
		return nil

	case `-d`, `-dir`:
		data["dir"] = args[1]
		return nil

	case `-p`, `-pkg`:
		data[PkgName] = args[1]
		return nil

	default:
		return fmt.Errorf("unknown command")
	}
}

func generateFile(data map[string]string) error {
	tmpl, err := template.New("gen-iter").Parse(iteratorTemplate)
	if err != nil {
		return err
	}

	if _, ok := data[Dir]; !ok {
		data[Dir] = defaultDir
	}

	f, err := os.Create(path.Join(data[Dir], fmt.Sprintf("%s_gen.go", data[IteratorName])))
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
