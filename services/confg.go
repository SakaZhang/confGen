package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type ConfgCsv struct {
	BaseStruct
}

func (u ConfgCsv) ToCSV() ([]string, error) {
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

func (u *ConfgCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("confg")
	lastConf, ok := tmp.(ConfgCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to ConfgCsv")
	}

	aux := &struct {
		Node []string `yaml:"node"`
		Dir  string   `yaml:"dir"`
		Pkg  string   `yaml:"pkg"`
		Env  string   `yaml:"env"`
	}{
		Node: u.Node,
		Dir:  lastConf.Dir,
		Pkg:  lastConf.Pkg,
		Env:  u.Env,
	}

	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML into auxiliary structure for UpCsv : %w", err)
	}

	u.Node = aux.Node
	u.Dir = aux.Dir
	u.Pkg = aux.Pkg
	u.Env = aux.Env
	// 校验非缺省字段
	validate := validator.New()
	err := validate.Struct(u)
	return err
}
