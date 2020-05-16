package strategy

//-----------------------------------------------
type Flyer interface {
	Fly() string
}

//-----------------------------------------------
type SlowFlyer struct{}

func (*SlowFlyer) Fly() string {
	return "Flying slow..."
}

//-----------------------------------------------
type FastFlyer struct{}

func (*FastFlyer) Fly() string {
	return "Flying so fast!!!"
}
