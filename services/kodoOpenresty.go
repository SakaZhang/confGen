package services

type KodoOpenrestyCsv struct {
	BaseStruct
	Idc  string `csv:"idc" yaml:"idc"`
	Tag  string `csv:"tag" yaml:"tag"`
	Sip  string `csv:"sip" yaml:"sip"`
	Sip2 string `csv:"sip2" yaml:"sip2"`
	Sip3 string `csv:"sip3" yaml:"sip3"`
}
