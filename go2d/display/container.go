package display

import (
	"github.com/wolfired/golabs/go2d/display/comps"
)

// CreateContainer 创建Container
func CreateContainer(name string) (c *Container) {
	c = &Container{Object: *CreateObject(name)}
	c.Parent = c.DryGet(comps.CCParent).(comps.Parent)
	return
}

// Container 显示容器
type Container struct {
	Object
	comps.Parent
}
