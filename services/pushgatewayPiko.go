package services

type PushgatewayPikoCsv struct {
	BaseStruct
	Data_dir string `csv:"data_dir" yaml:"data_dir"`
}
