package middleware

/*
Middleware 中间件
*/
type Middleware func(s *Context)

/*
Application 应用
*/
type Application struct {
	length      int
	middlewares [1024]Middleware
}

/*
Use 注册中间件
*/
func (a *Application) Use(m Middleware) {
	a.middlewares[a.length] = func(s *Context) {
		done := <-s.ready
		m(s)
		done <- struct{}{}
	}
	a.length++
}

/*
Handle 处理
*/
func (a *Application) Handle() {
	c := &Context{-1, make(chan chan struct{}), a}
	c.Next()
}

/*
Context 上下文
*/
type Context struct {
	index int
	ready chan chan struct{}
	m     *Application
}

/*
Next 跳转
*/
func (c *Context) Next() {
	c.index++
	if c.index == c.m.length {
		return
	}

	go c.m.middlewares[c.index](c)

	done := make(chan struct{})
	c.ready <- done
	<-done
}
