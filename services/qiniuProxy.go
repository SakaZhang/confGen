package services

type QiniuProxyCsv struct {
	BaseStruct
	Tag         string `csv:"tag" yaml:"tag"`
	Cache_disk  string `csv:"cache_disk" yaml:"cache_disk"`
	Cache_disk2 string `csv:"cache_disk2" yaml:"cache_disk2"`
	Cache_disk3 string `csv:"cache_disk3" yaml:"cache_disk3"`
	Cache_disk4 string `csv:"cache_disk4" yaml:"cache_disk4"`
	Cache_disk5 string `csv:"cache_disk5" yaml:"cache_disk5"`
	Cache_disk6 string `csv:"cache_disk6" yaml:"cache_disk6"`
}
