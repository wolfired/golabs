package comps

import (
	"github.com/wolfired/golabs/go2d/ecs"
)

var (
	// CCTransform Transform组件生成器
	CCTransform = ecs.Register(func(host ecs.IEntity) ecs.IComponent {
		return &transform{IComponent: ecs.CreateComponent("Transform", host)}
	})
)

// Transform 组件
type Transform interface {
	ecs.IComponent
}

type transform struct {
	ecs.IComponent
}

// Exec 业务
func (t *transform) Exec() {
}
