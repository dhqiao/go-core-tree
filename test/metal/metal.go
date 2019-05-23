package metal

type Imetal interface {
	GetName() string
	SetName(string) string
}

type Metal struct {
	Name string
	Exchange string
}

func (self Metal) GetName() string {
	if self.Name==""{
		return "none"
	}
	return self.Name
}

func (self *Metal) SetName(brand string) string {
	self.Name=brand
	return "done"
}