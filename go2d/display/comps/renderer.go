package comps

import (
	"github.com/wolfired/golabs/go2d/ecs"
)

var (
	// CCRenderer Renderer组件生成器
	CCRenderer = ecs.Register(func(host ecs.IEntity) ecs.IComponent {
		return &renderer{IComponent: ecs.CreateComponent("Renderer", host)}
	})
)

// Renderer 组件
type Renderer interface {
	ecs.IComponent
}

type renderer struct {
	ecs.IComponent
}

// Exec 业务
func (r *renderer) Exec() {
}
