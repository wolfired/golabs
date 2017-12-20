package idiotDB

const (
	FIELD_TYPE_INT    string = "int"
	FIELD_TYPE_UINT   string = "uint"
	FIELD_TYPE_STRING string = "string"
)

type MetaTable struct {
	Name   string
	Fields map[string]MetaField
}

func (t *MetaTable) RawData() []byte {
	zw := CreateZipWrapper()

	for fieldName, field := range t.Fields {
		zw.AddZipItem(fieldName+".zip", field.RawData())
	}

	zw.Close()

	return zw.RawData()
}

type MetaField struct {
	Name    string
	Type    string
	Default string
}

func (f *MetaField) RawData() []byte {
	zw := CreateZipWrapper()
	zw.AddZipItem("name", []byte(f.Name))
	zw.AddZipItem("type", []byte(f.Type))
	zw.AddZipItem("default", []byte(f.Default))
	zw.Close()

	return zw.RawData()
}
