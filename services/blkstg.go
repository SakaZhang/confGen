package services

type BlkstgCsv struct {
	BaseStruct
	Disk_count   int  `csv:"disk_count" yaml:"disk_count"`
	Disable_sync bool `csv:"disable_sync" yaml:"disable_sync"`
}
