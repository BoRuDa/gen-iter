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

var (
	ErrShowUsage       = fmt.Errorf("show usage")
	ErrPkgNameIsNotSet = fmt.Errorf("package name is not set")
	ErrTypeIsNotSet    = fmt.Errorf("type is not set")
	ErrUnknownFlag     = fmt.Errorf("unknown flag")
)

func main() {
	data, err := fetchDataFromArgs(os.Args[1:])
	if err != nil {
		handleErr(err)
		return
	}

	if err := generateFile(data); err != nil {
		handleErr(err)
		return
	}
}

func fetchDataFromArgs(args []string) (map[string]string, error) {
	data := make(map[string]string)

	if len(args) == 0 {
		return nil, ErrShowUsage
	}

	for validateArgs(args) {
		if err := setData(args[:2], data); err != nil {
			return nil, err
		}

		args = args[2:]
	}

	_, ok := data[PkgName]
	if !ok {
		return nil, ErrPkgNameIsNotSet
	}

	iterType, ok := data[Type]
	if !ok {
		return nil, ErrTypeIsNotSet
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
		return ErrUnknownFlag
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

func handleErr(err error) {
	fmt.Println(usage)
	if err != ErrShowUsage {
		fmt.Println("Error: ", err)
	}
}
