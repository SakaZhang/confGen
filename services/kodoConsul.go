package services

type KodoConsulCsv struct {
	BaseStruct
	Isserver     int    `csv:"isserver" yaml:"isserver"`
	Bind         string `csv:"bind" yaml:"bind"`
	Servernum    int    `csv:"servernum" yaml:"servernum"`
	Advertisewan int    `csv:"advertisewan" yaml:"advertisewan"`
}
