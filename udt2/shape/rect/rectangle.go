package rect

type Rect struct { // incase Rect does not start with Uppercase then it is also private but here it starts with Uppercase
	//length    float32 // length is private becase it starts with lowercase
	//bredth    float32 // bredth is private becase it starts with lowercase
	Length, Bredth float32 // Now these are public
	Area           float32
	Perimeter      float32
}

func AreaOfRect(r Rect) float32 {
	r.Area = r.Length * r.Bredth
	return r.Area
}
func PerimeterOfRect(r Rect) float32 {
	r.Perimeter = 2 * (r.Length + r.Bredth)
	return r.Perimeter
}
func AreaOfRectPtr(r *Rect) float32 {
	r.Area = r.Length * r.Bredth
	return r.Area
}
func PerimeterOfRectPtr(r *Rect) float32 {
	r.Perimeter = 2 * (r.Length + r.Bredth)
	return r.Perimeter
}

//this is called  a method .r is a receiver
func (r *Rect) AreaOf() float32 {
	r.Area = r.Length * r.Bredth
	return r.Area
}
func (r *Rect) PerimeterOf() float32 {
	r.Perimeter = 2 * (r.Length + r.Bredth)
	return r.Perimeter
}
func (r *Rect) ADD(a, b int) int {
	return a + b

}
