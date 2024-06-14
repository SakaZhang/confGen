package services

type KodoRedisClusterCsv struct {
	BaseStruct
	Port  int `csv:"port" yaml:"port"`
	Eport int `csv:"eport" yaml:"eport"`
}
