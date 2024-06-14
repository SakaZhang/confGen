package services

type KodoNsqdCsv struct {
	BaseStruct
	Tag           string `csv:"tag" yaml:"tag"`
	Iplook1       string `csv:"iplook1" yaml:"iplook1"`
	Iplook2       string `csv:"iplook2" yaml:"iplook2"`
	Tcp_port      int    `csv:"tcp_port" yaml:"tcp_port"`
	Http_port     int    `csv:"http_port" yaml:"http_port"`
	Data_path     string `csv:"data_path" yaml:"data_path"`
	Exporter_port int    `csv:"exporter_port" yaml:"exporter_port"`
}
