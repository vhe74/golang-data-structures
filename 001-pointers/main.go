package main

import "fmt"

func main() {
	var a int

	a = 2
	println("main", &a, " -> ", a)
	modifVal(a)
	println("main", &a, " -> ", a)
	modifRef(&a)
	println("main", &a, " -> ", a)

	//-----
	println("--------------------")
	buggyLoop()
	println("--------------------")
	fixedLoop()
}

func modifVal(v int) {

	println("modifVal", &v, " -> ", v)
	v = 3
	println("modifVal", &v, " -> ", v)
}

func modifRef(v *int) {
	println("modifRef", &v, " -> ", *v)
	*v = 4
	println("modifRef", &v, " -> ", *v)
}

//-----
type Dog struct {
	name string
}

func buggyLoop() {
	dogs := []Dog{Dog{"Ghost"}, Dog{"Bruno"}, Dog{"Lucky"}}
	dogPtrs := []*Dog{}
	for _, dog := range dogs {
		fmt.Printf("Dog with name <%s> and pointer: <%p>\n", dog, &dog)
		dogPtrs = append(dogPtrs, &dog)
	}

	for _, dogPtr := range dogPtrs {
		fmt.Printf("dogPtr with name <%s> and pointer: <%p>\n", dogPtr.name, dogPtr)
	}
}

func fixedLoop() {
	dogs := []Dog{Dog{"Ghost"}, Dog{"Bruno"}, Dog{"Lucky"}}
	dogPtrs := []*Dog{}
	for _, e := range dogs {
		e := e // use a local copy of e to avoid the recycling of the original e
		fmt.Printf("Dog with name <%s> and pointer: <%p>\n", e, &e)
		dogPtrs = append(dogPtrs, &e)
	}

	for _, p := range dogPtrs {
		fmt.Printf("dogPtr with name <%s> and pointer: <%p>\n", p.name, p)
	}
}
