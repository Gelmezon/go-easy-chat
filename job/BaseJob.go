package job

import "reflect"

type BaseJob interface {
	Execute()
	Register() reflect.Type
}

var TypeRegistry map[string]reflect.Type

func init() {
	TypeRegistry = make(map[string]reflect.Type)
}
