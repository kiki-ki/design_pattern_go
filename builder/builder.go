package builder

type builder interface {
  A() string
}

type Director struct {
	builder
}

func (d *Director) Construct() string {
	return d.A()
}

type ConcreateBuilder1 struct {}

func (c *ConcreateBuilder1) A() string {
	return "1"
}

type ConcreateBuilder2 struct {}

func (c *ConcreateBuilder2) A() string {
	return "2"
}
