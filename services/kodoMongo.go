package services

type KodoMongoCsv struct {
	BaseStruct
	Port            int    `csv:"port" yaml:"port"`
	Exporterport    int    `csv:"exporterport" yaml:"exporterport"`
	Version         string `csv:"version" yaml:"version"`
	Replset         string `csv:"replset" yaml:"replset"`
	OplogSize       string `csv:"oplogSize" yaml:"oplogSize"`
	Dbpath          string `csv:"dbpath" yaml:"dbpath"`
	Directoryperdb  string `csv:"directoryperdb" yaml:"directoryperdb"`
	StorageEngine   string `csv:"storageEngine" yaml:"storageEngine"`
	CacheSizeGB     int    `csv:"cacheSizeGB" yaml:"cacheSizeGB"`
	Auth            string `csv:"auth" yaml:"auth"`
	Master          string `csv:"master" yaml:"master"`
	Backup          string `csv:"backup" yaml:"backup"`
	Arbiter         string `csv:"arbiter" yaml:"arbiter"`
	BackupStatus    string `csv:"backupStatus" yaml:"backupStatus"`
	Script          string `csv:"script" yaml:"script"`
	BackupDisk      string `csv:"backupDisk" yaml:"backupDisk"`
	RsyncAddress    string `csv:"rsyncAddress" yaml:"rsyncAddress"`
	RsyncKey        string `csv:"rsyncKey" yaml:"rsyncKey"`
	RsyncPath       string `csv:"rsyncPath" yaml:"rsyncPath"`
	Tag             string `csv:"tag" yaml:"tag"`
	IgnoreNoopCheck string `csv:"ignoreNoopCheck" yaml:"ignoreNoopCheck"`
}
