package main 

import ("fmt"; "math")

func main(){
  fmt.Println(distance(0,0,1,1))
  r := Rectangle{x1: 0, y1: 0, x2: 2, y2: 2}
  fmt.Println(r.area())
  c := Circle{r: 2}
  fmt.Println(c.area())
  fmt.Println(r.perimeter())
  fmt.Println(c.perimeter())
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}



type Shape interface {
  area() float64
  perimeter() float64
}



type Rectangle struct{
  x1, y1, x2, y2 float64
}

func(r *Rectangle) length() float64{
  return distance(r.x1, r.y1, r.x2, r.y1)
}

func(r *Rectangle) width() float64{
  return distance(r.x1, r.y1, r.x1, r.y2)
}

func (r *Rectangle) area() float64{
  return r.length() * r.width()
}

func (r *Rectangle) perimeter() float64{
  return 2*r.length() + 2*r.width()
}



type Circle struct{
  x, y, r float64
}

func (c *Circle) area() float64{
  return math.Pi * (c.r*c.r)
}

func (c *Circle) perimeter() float64{
  return math.Pi * (c.r)
}