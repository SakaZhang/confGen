package services

type S3apiv2Csv struct {
	BaseStruct
	Tag        string `csv:"tag" yaml:"tag"`
	Cache_disk string `csv:"cache_disk" yaml:"cache_disk"`
	Idc        string `csv:"idc" yaml:"idc"`
	Proxy3     string `csv:"proxy3" yaml:"proxy3"`
}
