package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type CsvWriteable interface {
	ToCSV() ([]string, error)
}

// getLastConf 根据不同服务获取对应csv中最后一行的配置进行补全
func getLastConf(service string) interface{} {
	file := GitDir + fmt.Sprintf("/floy/%s/%s.csv", service, service)
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var lastRow []string
	var header []string
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		lastRow = row
		if header == nil {
			header = row
			continue
		}
	}

	headerMap := parseCSVHeader(header)
	// 这里比较麻烦, 因为作用域的问题, 写的很丑陋
	switch service {
	case "apiserver":
		var lastConf ApiServerCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("apiserver服务获取配置失败, 检查apiserver.csv.\n%v", err)
		}
		return lastConf
	case "up":
		var lastConf UpCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("up服务获取配置失败, 检查up.csv.\n%v", err)
		}
		return lastConf
	case "io":
		var lastConf IoCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("io服务获取配置失败, 检查up.csv.\n%v", err)
		}
		return lastConf
	case "bdchecker":
		var lastConf BdCheckerCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("bdchecker服务获取配置失败, 检查bdchecker.csv.\n%v", err)
		}
		return lastConf
	case "blackbox_exporter":
		var lastConf BlackboxExporterCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("blackbox_exporter服务获取配置失败, 检查blackbox_exporter.csv.\n%v", err)
		}
		return lastConf
	case "blkmaster":
		var lastConf BlkmasterCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("blkmaster服务获取配置失败, 检查blkmaster.csv.\n%v", err)
		}
		return lastConf
	case "blkmaster_exporter":
		var lastConf BlkmasterExporterCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("blkmaster_exporter服务获取配置失败, 检查blkmaster_exporter.csv.\n%v", err)
		}
		return lastConf
	case "blkstg":
		var lastConf BlkstgCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("blkstg服务获取配置失败, 检查blkstg.csv.\n%v", err)
		}
		return lastConf
	case "blkworker":
		var lastConf BlkworkerCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("blkworker服务获取配置失败, 检查blkworker.csv.\n%v", err)
		}
		return lastConf
	case "confg":
		var lastConf ConfgCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("confg服务获取配置失败, 检查confg.csv.\n%v", err)
		}
		return lastConf
	case "iorefresh":
		var lastConf IorefreshCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("iorefresh服务获取配置失败, 检查iorefresh.csv.\n%v", err)
		}
		return lastConf
	case "kodobill":
		var lastConf KodoBillCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("kodoBill服务获取配置失败, 检查kodoBill.csv.\n%v", err)
		}
		return lastConf
	case "kodo-consul":
		var lastConf KodoConsulCsv
		if err := csvToStruct(&lastConf, headerMap, lastRow); err != nil {
			log.Fatalf("kodo-consul服务获取配置失败, 检查kodo-consul.csv.\n%v", err)
		}
		return lastConf
	}

	return nil
}

// 工具函数, 解析csv文件的header
func parseCSVHeader(header []string) map[string]int {
	headerMap := make(map[string]int)
	for i, h := range header {
		headerMap[h] = i
	}

	return headerMap
}

// 工具函数, 将csv中的数据填充到对应服务结构体中
func csvToStruct(v interface{}, headerMap map[string]int, record []string) error {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		if field.Kind() == reflect.Struct {
			if err := csvToStruct(field.Addr().Interface(), headerMap, record); err != nil {
				return err
			}
			continue
		}

		tag := fieldType.Tag.Get("csv")

		if tag == "" {
			continue
		}

		if index, ok := headerMap[tag]; ok && index < len(record) {
			value := record[index]

			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Slice:
				if field.Type().Elem().Kind() == reflect.String {
					field.Set(reflect.ValueOf(strings.Split(value, ",")))
				}
			case reflect.Int:
				intValue, err := strconv.Atoi(value)
				if err != nil {
					return fmt.Errorf("error converting to int: %v", err)
				}
				field.SetInt(int64(intValue))
			case reflect.Float64:
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return fmt.Errorf("error converting to float64: %v", err)
				}
				field.SetFloat(floatValue)
			// Add more cases as needed for other types
			default:
				return fmt.Errorf("unsupported kind %s", field.Kind())
			}
		}
	}

	return nil
}

func structToCSV(v interface{}) (string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typeOfVal := val.Type()

	var record []string

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typeOfVal.Field(i)

		if field.Kind() == reflect.Struct && field.CanAddr() {
			subRecord, err := structToCSV(field.Addr().Interface())
			if err != nil {
				return "", err
			}
			record = append(record, strings.Split(subRecord, ",")...)
			continue
		}

		tag := fieldType.Tag.Get("csv")
		if tag == "" || !field.CanInterface() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			record = append(record, field.String())
		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.String {
				record = append(record, strings.Join(field.Interface().([]string), ","))
			}
		case reflect.Int:
			record = append(record, strconv.Itoa(int(field.Int())))
		case reflect.Float64:
			record = append(record, strconv.FormatFloat(field.Float(), 'f', -1, 64))
		default:
			return "", fmt.Errorf("unsupported kind %s", field.Kind())
		}
	}

	return strings.Join(record, ","), nil
}
