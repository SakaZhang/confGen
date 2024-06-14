package services

type RollbackCheckCsv struct {
	BaseStruct
	Tag string `csv:"tag" yaml:"tag"`
	Id  string `csv:"id" yaml:"id"`
}
