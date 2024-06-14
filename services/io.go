package services

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type IoCsv struct {
	BaseStruct
	Tag        string `csv:"tag" yaml:"tag"`
	Cache_disk string `csv:"cache_disk" yaml:"cache_disk"`
	Idc        string `csv:"idc" yaml:"idc"`
}

func (u IoCsv) ToCSV() ([]string, error) {
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

func (u *IoCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("io")
	lastConf, ok := tmp.(IoCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to UpCsv")
	}

	type Alias IoCsv // Create an alias to avoid recursion
	aux := &struct {
		Node      []string `yaml:"node"`
		Dir       string   `yaml:"dir"`
		Pkg       string   `yaml:"pkg"`
		Env       string   `yaml:"env"`
		Tag       string   `yaml:"tag"`
		CacheDisk string   `yaml:"cache_disk"` // 使用驼峰命名法
		Idc       string   `yaml:"idc"`
	}{
		Node:      u.Node,
		Dir:       lastConf.Dir,
		Pkg:       lastConf.Pkg,
		Env:       u.Env,
		Tag:       u.Tag,
		CacheDisk: lastConf.Cache_disk,
		Idc:       u.Idc,
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
	if aux.Idc == "" {
		u.Idc = aux.Env
	} else {
		u.Idc = aux.Idc
	}

	return nil
}
