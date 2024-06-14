package services

type PrometheusPikoCsv struct {
	BaseStruct
	Port                 int    `csv:"port" yaml:"port"`
	Retention            int    `csv:"retention" yaml:"retention"`
	Replica_member_total int    `csv:"replica_member_total" yaml:"replica_member_total"`
	Replica_member_num   int    `csv:"replica_member_num" yaml:"replica_member_num"`
	Friendly_name        string `csv:"friendly_name" yaml:"friendly_name"`
	Region               string `csv:"region" yaml:"region"`
	Is_shard             int    `csv:"is_shard" yaml:"is_shard"`
	Replica_num          int    `csv:"replica_num" yaml:"replica_num"`
	Bdir                 string `csv:"bdir" yaml:"bdir"`
	Is_dynamic           string `csv:"is_dynamic" yaml:"is_dynamic"`
}
