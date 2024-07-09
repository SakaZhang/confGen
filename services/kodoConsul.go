package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
)

type KodoConsulCsv struct {
	BaseStruct
	Isserver     []string `csv:"isserver" yaml:"isserver" validate:"required"`
	Bind         string   `csv:"bind" yaml:"bind"`
	Servernum    string   `csv:"servernum" yaml:"servernum"`
	Advertisewan string   `csv:"advertisewan" yaml:"advertisewan"`
}

func handle(node string, m map[string]bool, u KodoConsulCsv) (string, error) {
	u.BaseStruct.Node = []string{node}
	u.Isserver = nil
	csvRecord, err := structToCSV(&u)
	if err != nil {
		panic(err)
	}
	// if node is server, here are some differences
	fmt.Println(csvRecord)
	split := strings.Split(csvRecord, ",")
	if ok := m[node]; ok {
		split[1] = "kodo-consul-server"
		split[4] = "1"
		split[5] = "0.0.0.0"
		split[6] = strconv.Itoa(len(m))
		// TODO get advertise_wan
		split[7] = "ADVERTISE_WAN"
		csvRecord = strings.Join(split, ",")
	} else {
		split[1] = "kodo-consul-agent"
		for i := 4; i < len(split); i++ {
			split[i] = ""
		}
		csvRecord = strings.Join(split, ",")
	}

	return strings.Join(split, ","), nil
}

func (u KodoConsulCsv) ToCSV() ([]string, error) {
	var csv []string
	m := make(map[string]bool)
	for _, v := range u.Isserver {
		m[strings.TrimSpace(v)] = true
	}
	for _, node := range u.Node {
		csvRecord, err := handle(node, m, u)
		if err != nil {
			return nil, err
		}
		csv = append(csv, csvRecord)
	}
	return csv, nil
}

func (u *KodoConsulCsv) UnmarshalYAML(value *yaml.Node) error {
	tmp := getLastConf("kodo-consul")
	lastConf, ok := tmp.(KodoConsulCsv)
	if !ok {
		return fmt.Errorf("failed to cast lastConf to KodoBillCsv")
	}

	aux := &struct {
		Node         []string `yaml:"node"`
		Dir          string   `yaml:"dir"`
		Pkg          string   `yaml:"pkg"`
		Env          string   `yaml:"env"`
		Isserver     []string `csv:"isserver" yaml:"isserver" validate:"required"`
		Bind         string   `csv:"bind" yaml:"bind"`
		Servernum    string   `csv:"servernum" yaml:"servernum"`
		Advertisewan string   `csv:"advertisewan" yaml:"advertisewan"`
	}{
		Node:         u.Node,
		Dir:          lastConf.Dir,
		Pkg:          lastConf.Pkg,
		Env:          u.Env,
		Isserver:     nil,
		Bind:         "",
		Servernum:    "",
		Advertisewan: "",
	}

	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML into auxiliary structure for kodo-consul : %w", err)
	}

	u.Node = aux.Node
	u.Dir = aux.Dir
	u.Pkg = aux.Pkg
	u.Env = aux.Env
	u.Isserver = aux.Isserver
	u.Bind = aux.Bind
	u.Servernum = aux.Servernum
	u.Advertisewan = aux.Advertisewan
	// 校验非缺省字段
	validate := validator.New()
	err := validate.Struct(u)
	return err
}
