package ecs

import (
	"fmt"
)

// ComponentCreater 组件生成器
type ComponentCreater *componentCreater

type componentCreater struct {
	create func(host IEntity) IComponent
}

// Register 注册组件生成器
func Register(creater func(host IEntity) IComponent) ComponentCreater {
	return &componentCreater{create: creater}
}

// IComponent 组件
type IComponent interface {
	// 返回组件唯一ID
	ID() uint
	// 返回组件名
	Name() string
	Host() IEntity
	Exec()
	lost()
}

// CreateComponent 创建默认组件
func CreateComponent(name string, host IEntity) IComponent {
	return &component{name: name, host: host}
}

type component struct {
	id   uint
	name string
	host IEntity
}

// ID 返回组件唯一ID
func (c *component) ID() uint {
	return c.id
}

// Name 返回组件名
func (c *component) Name() string {
	return c.name
}

// Host 返回宿主实体
func (c *component) Host() IEntity {
	return c.host
}

// Exec 默认组件业务逻辑
func (c *component) Exec() {
	fmt.Println("default Component.Exec()")
}

// lost (从父实体删除后)清除父实体
func (c *component) lost() {
	c.host = nil
}
