package display

import (
	"github.com/wolfired/golabs/go2d/display/comps"
	"github.com/wolfired/golabs/go2d/ecs"
)

// CreateObject 创建Object
func CreateObject(name string) (o *Object) {
	o = &Object{IEntity: ecs.CreateEntity(name)}
	o.Child = o.DryGet(comps.CCChild).(comps.Child)
	return
}

// Object 显示对象
type Object struct {
	ecs.IEntity
	comps.Child
}
