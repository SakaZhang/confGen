package services

var GitDir string

type BaseStruct struct {
	Node []string `csv:"node" yaml:"node" binding:"required"`
	Dir  string   `csv:"dir"  yaml:"dir"`
	Pkg  string   `csv:"pkg"  yaml:"pkg"`
	Env  string   `csv:"env"  yaml:"env" binding:"required"`
}
