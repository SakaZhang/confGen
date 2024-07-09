package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"strings"
)

type BlkmasterCsv struct {
	BaseStruct
	IS_PRIMARY string `csv:"IS_PRIMARY" yaml:"IS_PRIMARY"`
}

func (u BlkmasterCsv) ToCSV() ([]string, error) {
	var csv []string
	for _, node := range u.Node {
		u.BaseStruct.Node = []string{node}
		csvRecord, err := structToCSV(&u)

		// set IS_PRIMARY to true if node is primary
		if node == u.IS_PRIMARY {
			split := strings.Split(csvRecord, ",")
			split[len(split)-1] = "true"
			csvRecord = strings.Join(split, ",")
		} else {
			split := strings.Split(csvRecord, ",")
			split[len(split)-1] = "false"
			csvRecord = strings.Join(split, ",")
		}

		if err != nil {
			panic(err)
		}
		csv = append(csv, csvRecord)
	}
	return csv, nil
}

func (u *BlkmasterCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("blkmaster")
	lastConf, ok := tmp.(BlkmasterCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to BlkmasterCsv")
	}

	type Alias IoCsv // Create an alias to avoid recursion
	aux := &struct {
		Node       []string `yaml:"node"`
		Dir        string   `yaml:"dir"`
		Pkg        string   `yaml:"pkg"`
		Env        string   `yaml:"env"`
		IS_PRIMARY string   `yaml:"IS_PRIMARY"`
	}{
		Node:       u.Node,
		Dir:        lastConf.Dir,
		Pkg:        lastConf.Pkg,
		Env:        u.Env,
		IS_PRIMARY: "false",
	}

	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML into auxiliary structure for UpCsv : %w", err)
	}
	u.Node = aux.Node
	u.Dir = aux.Dir
	u.Pkg = aux.Pkg
	u.Env = aux.Env
	u.IS_PRIMARY = aux.IS_PRIMARY

	validate := validator.New()
	err := validate.Struct(u)

	return err
}
