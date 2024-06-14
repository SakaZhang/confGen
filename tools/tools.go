package tools

import (
	"confGenerator/services"
	"fmt"
)

// Config 存储配置文件中的配置信息并结合各服务中csv定义进行补全, 其中node字段为[]string类型
type Config struct {
	Apiserver services.ApiServerCsv `yaml:"apiserver" csv:"apiserver.csv"`
	Up        services.UpCsv        `yaml:"up" csv:"up.csv"`
	Io        services.IoCsv        `yaml:"io" csv:"io.csv"`
}

func GenConfFromCfgFile(cfgFile string) error {
	// 补全配置文件中的配置信息
	cfg, err := compConfWithCfgFile(cfgFile)
	if err != nil {
		return fmt.Errorf("补全配置失败: %v", err)
	}

	err = writeCfgIntoCsv(cfg)
	if err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	return nil
}
