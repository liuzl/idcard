package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

var (
	path = `src/github.com/liuzl/idcard`
)

func genCode(m map[string]string, file string) {
	output := bytes.Buffer{}
	output.WriteString("package idcard\n\n")
	output.WriteString("var CodeMap = map[string]string{\n")
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		output.WriteString(fmt.Sprintf("\t\"%s\": \"%s\",\n", k, m[k]))
	}
	output.WriteString("}")
	if err := ioutil.WriteFile(file, output.Bytes(), os.FileMode(0664)); err != nil {
		log.Fatal(err)
	}
}

func main() {
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("Missing $GOPATH environment variable")
	}
	files, err := filepath.Glob(filepath.Join(gopath, path) + "/data/*.json")
	if err != nil {
		log.Fatal(err)
	}
	ret := make(map[string]string)
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		var m map[string]string
		err = json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range m {
			ret[k] = v
		}
	}
	file := filepath.Join(gopath, path, "codemap.go")
	genCode(ret, file)
}
