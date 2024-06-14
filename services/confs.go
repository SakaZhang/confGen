package services

type ConfsCsv struct {
	BaseStruct
	Tag      string `csv:"tag" yaml:"tag"`
	Is_proxy bool   `csv:"is_proxy" yaml:"is_proxy"`
}
