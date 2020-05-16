package strategy

//-----------------------------------------------
type Duck struct {
	flyer   Flyer
	quacker Quacker
}

func NewDuck(f Flyer, q Quacker) *Duck {
	return &Duck{flyer: f, quacker: q}
}

func (*Duck) Swim() string {
	return "swimming..."
}

func (d *Duck) Fly() string {
	return d.flyer.Fly()
}

func (d *Duck) Quack() string {
	return d.quacker.Quack()
}

func (d *Duck) Display() string {
	return "I am a duck"
}

//-----------------------------------------------
type SlowQuietDuck struct {
	*Duck
}

func NewSlowQuietDuck() *SlowQuietDuck {
	return &SlowQuietDuck{NewDuck(&SlowFlyer{}, &QuietQuacker{})}
}

func (*SlowQuietDuck) Sleep() string {
	return "sleeping..."
}

func (*SlowQuietDuck) Display() string {
	return "I am a slow and quiet duck..."
}

//-----------------------------------------------
type FastLoudDuck struct {
	*Duck
}

func NewFastLoudDuck() *FastLoudDuck {
	return &FastLoudDuck{NewDuck(&FastFlyer{}, &LoudQuacker{})}
}

func (*FastLoudDuck) Run() string {
	return "Running!!"
}

func (*FastLoudDuck) Display() string {
	return "I am a fast and loud duck!!!"
}
