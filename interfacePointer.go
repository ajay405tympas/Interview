// You can edit this code!
// Click here and start typing.
package main

type A interface {
	Apple()
}

type B interface {
	Ball()
}

type C interface {
	Apple()
	Ball()
}

type X struct{}

func (*X) Apple() { println("A printed") }

func (X) Ball() { println("B printed") }

func main() {

	var c C = X{}
	c.Ball()

	//c := X{}

	//c.Apple()

}
