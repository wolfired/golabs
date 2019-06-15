package ecs

// IEntity 实体
type IEntity interface {
	// 返回实体唯一ID
	ID() uint
	// 返回实体名
	Name() string
	TryGet(cc ComponentCreater) IComponent
	DryGet(cc ComponentCreater) IComponent
	Delete(cc ComponentCreater) IComponent
}

// CreateEntity 创建实体
func CreateEntity(name string) IEntity {
	return &entity{name: name, comps: make(map[ComponentCreater]IComponent)}
}

type entity struct {
	id    uint
	name  string
	comps map[ComponentCreater]IComponent
}

// ID 返回实体唯一ID
func (e *entity) ID() uint {
	return e.id
}

// Name 返回实体名
func (e *entity) Name() string {
	return e.name
}

// TryGet (如果存在则)返回组件
func (e *entity) TryGet(cc ComponentCreater) IComponent {
	c := e.comps[cc]
	return c
}

// DryGet (如果不存在则创建并)返回组件
func (e *entity) DryGet(cc ComponentCreater) IComponent {
	c := e.comps[cc]

	if nil == c {
		c = cc.create(e)
		e.comps[cc] = c
	}

	return c
}

// Del (如果存在则)删除并返回组件
func (e *entity) Delete(cc ComponentCreater) IComponent {
	c := e.comps[cc]

	if nil != c {
		delete(e.comps, cc)
		c.lost()
	}

	return c
}
