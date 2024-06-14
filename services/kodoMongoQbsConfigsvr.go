package services

type KodoMongoQbsConfigsvrCsv struct {
	BaseStruct
	Configsvr_sharding_stat string `csv:"configsvr_sharding_stat" yaml:"configsvr_sharding_stat"`
	KodoMongoCsv
}
