package services

type PikoMongoExporterCsv struct {
	BaseStruct
	Port                    string `csv:"port" yaml:"port"`
	Mgoport                 string `csv:"mgoport" yaml:"mgoport"`
	Configsvr_sharding_stat string `csv:"configsvr_sharding_stat" yaml:"configsvr_sharding_stat"`
	Enable_dbstats          string `csv:"enable_dbstats" yaml:"enable_dbstats"`
}
