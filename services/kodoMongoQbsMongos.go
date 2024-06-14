package services

type KodoMongoQbsMongosCsv struct {
	BaseStruct
	Fail_sleep_ms int    `csv:"fail_sleep_ms" yaml:"fail_sleep_ms"`
	Port          int    `csv:"port" yaml:"port"`
	Exporterport  int    `csv:"exporterport" yaml:"exporterport"`
	Configdb      string `csv:"configdb" yaml:"configdb"`
	Auth          string `csv:"auth" yaml:"auth"`
}
