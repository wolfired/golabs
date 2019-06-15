package comps

import (
	"github.com/wolfired/golabs/go2d/ecs"
)

var (
	// CCChild Child组件生成器
	CCChild = ecs.Register(func(host ecs.IEntity) ecs.IComponent {
		return &child{IComponent: ecs.CreateComponent("Child", host)}
	})
)

// Child 组件
type Child interface {
	ecs.IComponent

	Parent() Parent
}

type child struct {
	ecs.IComponent

	order  int
	parent Parent
}

func (c *child) Parent() Parent {
	return c.parent
}
