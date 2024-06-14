package tools

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func compConfWithCfgFile(cfgFile string) (Config, error) {
	f, err := os.Open(cfgFile)
	if err != nil {
		return Config{}, fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer f.Close()

	var cfg Config
	// 为每个服务实现了UnmarshalYAML方法, 在解析时就补全config.yaml中缺省的信息
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return cfg, nil
}
