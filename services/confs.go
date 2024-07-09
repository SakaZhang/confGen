package services

/*type ConfsCsv struct {
	BaseStruct
	Tag      string `csv:"tag" yaml:"tag"`
	Is_proxy bool   `csv:"is_proxy" yaml:"is_proxy"`
}
*/

type ConfsCsv struct {
	Dir1
	Dir2
}

type Dir1 struct {
	BaseStruct
	Tag      string `csv:"tag" yaml:"tag"`
	Is_proxy bool   `csv:"is_proxy" yaml:"is_proxy"`
}

type Dir2 struct {
	BaseStruct
	Tag      string `csv:"tag" yaml:"tag"`
	Is_proxy bool   `csv:"is_proxy" yaml:"is_proxy"`
}

// TODO
