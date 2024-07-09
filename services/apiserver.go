package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type ApiServerCsv struct {
	BaseStruct
	Tag string `csv:"tag" yaml:"tag"`
}

func (u ApiServerCsv) ToCSV() ([]string, error) {
	var csv []string
	for _, node := range u.Node {
		u.BaseStruct.Node = []string{node}
		csvRecord, err := structToCSV(&u)
		if err != nil {
			panic(err)
		}
		csv = append(csv, csvRecord)
	}
	return csv, nil
}

func (u *ApiServerCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("apiserver")
	lastConf, ok := tmp.(ApiServerCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to apiserver.")
	}

	aux := &struct {
		Node []string `yaml:"node"`
		Dir  string   `yaml:"dir"`
		Pkg  string   `yaml:"pkg"`
		Env  string   `yaml:"env"`
		Tag  string   `yaml:"tag"`
	}{
		Node: u.Node,
		Dir:  lastConf.Dir,
		Pkg:  lastConf.Pkg,
		Env:  u.Env,
		Tag:  u.Tag,
	}

	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML into auxiliary structure for UpCsv : %w", err)
	}
	u.Node = aux.Node
	u.Dir = aux.Dir
	u.Pkg = aux.Pkg
	u.Env = aux.Env
	u.Tag = aux.Tag

	validate := validator.New()
	err := validate.Struct(u)

	return err
}
