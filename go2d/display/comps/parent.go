package comps

import (
	"fmt"

	"github.com/wolfired/golabs/go2d/ecs"
)

var (
	// CCParent Parent组件生成器
	CCParent = ecs.Register(func(host ecs.IEntity) ecs.IComponent {
		return &parent{IComponent: ecs.CreateComponent("Parent", host)}
	})
)

// Parent 组件
type Parent interface {
	ecs.IComponent

	Add(child Child)
	Del(order int)
	At(order int)
}

type parent struct {
	ecs.IComponent
}

func (p *parent) Add(child Child) {
	fmt.Println("add child")
}

func (p *parent) Del(order int) {
	fmt.Println("del child at", order)
}

func (p *parent) At(order int) {
	fmt.Println("child at", order)
}
