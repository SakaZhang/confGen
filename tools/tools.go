package tools

import (
	"confGenerator/services"
	"encoding/csv"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
	"strings"
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
		log.Fatalf("补全配置失败!\nErr: %v", err)
	}

	err = writeCfgIntoCsv(cfg)
	if err != nil {
		log.Fatalf("写入配置失败!\nErr: %v", err)
	}

	return nil
}

func compConfWithCfgFile(cfgFile string) (Config, error) {
	f, err := os.Open(cfgFile)
	if err != nil {
		log.Fatalf("打开配置文件失败!")
		return Config{}, err
	}
	defer f.Close()
	var cfg Config
	// 因为为每个服务实现了UnmarshalYAML方法, 所以在解析时就补全了config.yaml中缺省的信息
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("解析配置文件失败!")
		return Config{}, err
	}
	// 现在cfg中已经补全了配置文件中的配置信息
	return cfg, nil
}

// TODO
func writeCfgIntoCsv(config Config) error {
	v := reflect.ValueOf(config)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()

		if csvWritable, ok := fieldValue.(services.CsvWriteable); ok {
			fieldType := t.Field(i)
			filenameTag := fieldType.Tag.Get("csv")
			filename := strings.TrimSpace(filenameTag)
			filepath := services.GitDir + "/floy/" + strings.ToLower(t.Field(i).Name) + "/" + filename
			if filepath == "" {
				return fmt.Errorf("field %s does not have a 'csv' tag", fieldType.Name)
			}

			if !strings.HasSuffix(filepath, ".csv") {
				filepath += ".csv"
			}

			if err := writeToCSV(filepath, csvWritable); err != nil {
				return fmt.Errorf("failed to write %s: %w", filename, err)
			}
		} else {
			continue
			log.Fatalf("field %s does not implement CsvWriteable", t.Field(i).Name)
		}
	}

	return nil
}

func writeToCSV(filepath string, data services.CsvWriteable) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	records, err := data.ToCSV()
	if err != nil {
		return err
	}
	// 插入一行空行
	err = writer.Write([]string{})
	if err != nil {
		return err
	}

	for _, record := range records {
		fmt.Println(record)
		err := writer.Write(strings.Split(record, ","))
		if err != nil {
			return err
		}
	}

	return nil
}
