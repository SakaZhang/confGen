package tools

import (
	"confGenerator/services"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

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
			return fmt.Errorf("field %s does not implement CsvWriteable", t.Field(i).Name)
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
		err := writer.Write(strings.Split(record, ","))
		if err != nil {
			return err
		}
	}

	return nil
}
