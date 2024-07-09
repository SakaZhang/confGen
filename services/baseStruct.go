package services

var GitDir string

type BaseStruct struct {
	Node []string `csv:"node" yaml:"node" validate:"required"`
	Dir  string   `csv:"dir"  yaml:"dir"`
	Pkg  string   `csv:"pkg"  yaml:"pkg"`
	Env  string   `csv:"env"  yaml:"env" validate:"required"`
}
