package iter

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestName2(t *testing.T) {
	tmpl, err := template.New("gen-iter").Parse(iteratorTemplate)
	t.Log(err)

	f, err := os.Create("./test/test_template.go")
	if err != nil {
		return
	}
	defer f.Close()

	name := "iter"
	iterType := "int"

	//TODO

	iteratorName := fmt.Sprint(strings.ToUpper(iterType[:1]), iterType[1:], "Iterator")

	err = tmpl.Execute(f, map[string]string{
		"PkgName":      strings.ToLower(name),
		"Type":         iterType,
		"IteratorName": iteratorName,
	})
	t.Log(err)
}
