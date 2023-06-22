package interfaceNest

// 接口嵌套
type Flyer interface {
	fly()
}
type Swimer interface {
	swim()
}
type FlyFish interface {
	Flyer
	Swimer
}
type Fish struct {
}

func (fish Fish) fly() {
	println("fly method run")
}
func (fish Fish) swim() {
	println("swim method run")
}

/** OCP设计原则实现 */
type Pet interface {
	eat()
	sleep()
}
type Dog struct {
}
type Cat struct {
}

func (dog Dog) eat() {
	println("dog eat...")
}
func (dog Dog) sleep() {
	println("dog sleep...")
}
func (cat Cat) eat() {
	println("cat eat...")
}
func (cat Cat) sleep() {
	println("cat sleep...")
}

type Person struct{}

// pet既可以传dog，也可以传递cat，因为dog和cat都是Pet类型
func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func InterfaceNest() {
	println("========================Interface Nest========================")
	println()

	// 接口嵌套
	var flyFish FlyFish
	flyFish = Fish{}
	flyFish.fly()
	flyFish.swim()

	/** OCP设计原则实现 */
	println()
	dog := Dog{}
	cat := Cat{}
	person := Person{}
	person.care(dog)
	person.care(cat)

	println()
	println("========================Interface Nest========================")
}
