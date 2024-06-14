package services

type ServiceConfigCsv struct {
	Node string `csv:"node" yaml:"node"`
	Dir  string `csv:"dir" yaml:"dir"`
	Env  string `csv:"env" yaml:"env"`
}
