package services

type KodoNsqlookupdCsv struct {
	BaseStruct
	Tag       string `csv:"tag" yaml:"tag"`
	Tcp_port  int    `csv:"tcp_port" yaml:"tcp_port"`
	Http_port int    `csv:"http_port" yaml:"http_port"`
}
