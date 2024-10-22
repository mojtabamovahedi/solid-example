package main

import "fmt"

func main() {

}

// bad example

//type Staff interface {
//	Develop()
//	WriteDocs()
//	Test()
//	Design()
//	Monitor()
//}
//
//type GolangJunior struct {
//	name string
//}
//
//func (g GolangJunior) Develop() {
//	fmt.Println("developing go...")
//}
//
//func (g GolangJunior) WriteDocs() {
//	fmt.Println("write weekly docs")
//}
//
//func (g GolangJunior) Test() {
//	panic("cant test")
//}
//
//func (g GolangJunior) Design() {
//	panic("cant design")
//}
//
//func (g GolangJunior) Monitor() {
//	panic("cant monitor")
//}
//
//type QATest struct {
//	name string
//}
//
//func (Q QATest) Develop() {
//	panic("i just know python")
//}
//
//func (Q QATest) WriteDocs() {
//	fmt.Println("write docs about their test")
//}
//
//func (Q QATest) Test() {
//	fmt.Println("know how to roast developer")
//}
//
//func (Q QATest) Design() {
//	panic("implement me")
//}
//
//func (Q QATest) Monitor() {
//	panic("implement me")
//}

// good example

type TesterStaff interface {
	Test()
	WriteDocs()
}

type DeveloperStaff interface {
	Develop()
}

type DesignerStaff interface {
	Design()
}

type DevOpsStaff interface {
	Monitor()
}

type GolangJunior struct {
	name string
	id   int
}

func (g GolangJunior) Develop() {
	fmt.Println("developing go...")
}

type QATester struct {
	name string
	id   int
}

func (Q QATester) Develop() {
	fmt.Println("i just know python")
}

func (Q QATester) WriteDocs() {
	fmt.Println("write docs about their test")
}

func (Q QATester) Test() {
	fmt.Println("know how to roast developer")
}
