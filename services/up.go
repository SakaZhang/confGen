package services

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type UpCsv struct {
	BaseStruct
	Tag        string `csv:"tag" yaml:"tag"`
	Cache_disk string `csv:"cache_disk" yaml:"cache_disk" binding:"required"`
	Idc        string `csv:"idc" yaml:"idc"`
}

func (u UpCsv) ToCSV() ([]string, error) {
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

func (u *UpCsv) UnmarshalYAML(value *yaml.Node) error {
	// Get last config in deploy/floy/SERVICE/SERVICE.csv as default value
	tmp := getLastConf("up")
	lastConf, ok := tmp.(UpCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to UpCsv")
	}

	type Alias UpCsv // Create an alias to avoid recursion
	aux := &struct {
		Node      []string `yaml:"node"`
		Dir       string   `yaml:"dir"`
		Pkg       string   `yaml:"pkg"`
		Env       string   `yaml:"env"`
		Tag       string   `yaml:"tag"`
		CacheDisk string   `yaml:"cache_disk"`
		Idc       string   `yaml:"idc"`
	}{
		Node:      u.Node,
		Dir:       lastConf.Dir,
		Pkg:       lastConf.Pkg,
		Env:       u.Env,
		Tag:       u.Tag,
		CacheDisk: u.Cache_disk,
		Idc: func() string {
			if u.Idc == "" {
				return u.Env
			}
			return u.Idc
		}(),
	}
	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML into auxiliary structure for UpCsv : %w", err)
	}

	u.Node = aux.Node
	u.Dir = aux.Dir
	u.Pkg = aux.Pkg
	u.Env = aux.Env
	u.Tag = aux.Tag
	u.Cache_disk = aux.CacheDisk
	u.Idc = aux.Idc

	return nil
}
