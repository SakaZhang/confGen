package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type BlkstgCsv struct {
	BaseStruct
	Disk_count   int    `csv:"disk_count" yaml:"disk_count" validate:"required"`
	Disable_sync string `csv:"disable_sync" yaml:"disable_sync" `
}

func (u BlkstgCsv) ToCSV() ([]string, error) {
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

func (u *BlkstgCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("blkstg")
	lastConf, ok := tmp.(BlkstgCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to BlkstgCsv")
	}

	aux := &struct {
		Node         []string `yaml:"node"`
		Dir          string   `yaml:"dir"`
		Pkg          string   `yaml:"pkg"`
		Env          string   `yaml:"env"`
		Disk_count   int      `yaml:"disk_count"`
		Disable_sync string   `yaml:"disable_sync"`
	}{
		Node:         u.Node,
		Dir:          lastConf.Dir,
		Pkg:          lastConf.Pkg,
		Env:          u.Env,
		Disk_count:   u.Disk_count,
		Disable_sync: "false",
	}

	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML into auxiliary structure for UpCsv : %w", err)
	}

	u.Node = aux.Node
	u.Dir = aux.Dir
	u.Pkg = aux.Pkg
	u.Env = aux.Env
	u.Disk_count = aux.Disk_count
	u.Disable_sync = aux.Disable_sync
	// 校验非缺省字段
	validate := validator.New()
	err := validate.Struct(u)
	return err
}
