package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type BlkmasterExporterCsv struct {
	BaseStruct
}

func (u BlkmasterExporterCsv) ToCSV() ([]string, error) {
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

func (u *BlkmasterExporterCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("blkmaster_exporter")
	lastConf, ok := tmp.(BlkmasterExporterCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to BlkmasterExporterCsv")
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

	validate := validator.New()
	err := validate.Struct(u)

	return err
}
