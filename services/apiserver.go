package services

type ApiServerCsv struct {
	BaseStruct
	Tag string `csv:"tag" yaml:"tag"`
}
