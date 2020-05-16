package strategy

//-----------------------------------------------
type Quacker interface {
	Quack() string
}

//-----------------------------------------------
type LoudQuacker struct{}

func (*LoudQuacker) Quack() string {
	return "QUACK!"
}

//-----------------------------------------------
type QuietQuacker struct{}

func (*QuietQuacker) Quack() string {
	return "qua..."
}
