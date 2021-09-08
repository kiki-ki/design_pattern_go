package factoryMethod

type Product interface {
	Use() string
}

type Factory interface {
	createProduct(owner string) Product
	registerProduct(p Product)
}

type factory struct {}

func (f *factory) Create(c Factory, owner string) Product {
	p := c.createProduct(owner)
	c.registerProduct(p)
	return p
}

type idCard struct {
	owner string
}

func (i *idCard) Use() string {
	return i.owner
}

type IDCardFactory struct {
	*factory
	owners []string
}

func (i *IDCardFactory) createProduct(owner string) Product {
	return &idCard{owner}
}

func (i *IDCardFactory) registerProduct(p Product) {
	i.owners = append(i.owners, p.Use())
}
