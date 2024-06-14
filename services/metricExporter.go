package services

type MetricExporterCsv struct {
	BaseStruct
	Idc string `csv:"idc" yaml:"idc"`
}
