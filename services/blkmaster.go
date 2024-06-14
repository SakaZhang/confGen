package services

type BlkmasterCsv struct {
	BaseStruct
	IS_PRIMARY bool `csv:"IS_PRIMARY" yaml:"IS_PRIMARY"`
}
