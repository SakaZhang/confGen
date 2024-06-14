package services

type KodoRedisQbsClusterCsv struct {
	KodoRedisClusterCsv
	Disk string `csv:"disk" yaml:"disk"`
}
