package services

type PfdDeleteCsv struct {
	BaseStruct
	Del_worker_num               int `csv:"del_worker_num" yaml:"del_worker_num"`
	Backup_del_worker_num        int `csv:"backup_del_worker_num" yaml:"backup_del_worker_num"`
	Multiregion_del_worker_num   int `csv:"multiregion_del_worker_num" yaml:"multiregion_del_worker_num"`
	Requeue_delay_ms             int `csv:"requeue_delay_ms" yaml:"requeue_delay_ms"`
	Backup_requeue_delay_ms      int `csv:"backup_requeue_delay_ms" yaml:"backup_requeue_delay_ms"`
	Multiregion_requeue_delay_ms int `csv:"multiregion_requeue_delay_ms" yaml:"multiregion_requeue_delay_ms"`
}
