package idiotDB

const (
	FIELD_TYPE_INT    string = "int"
	FIELD_TYPE_UINT   string = "uint"
	FIELD_TYPE_STRING string = "string"
)

type Meta_Table struct {
	Name   string
	Fields map[string]Meta_Field
}

func (self *Meta_Table) toZip() {

}

type Meta_Field struct {
	Name    string
	Type    string
	Default string
}

func (self *Meta_Field) toZip() {

}
