package services

type NsqExporterCsv struct {
	BaseStruct
	Tag string `csv:"tag" yaml:"tag"`
}
