package main

func main() {

}

type OrderManager struct {
}

// bad method

func (o *OrderManager) Process() {
	o.logOrder()
	o.sendOrder()
	o.saveOrder()
}

func (o *OrderManager) saveOrder() {}
func (o *OrderManager) sendOrder() {}
func (o *OrderManager) logOrder()  {}

// good way
//
//func (o *OrderManager) Process() {
//
//	log.Println(o)
//
//	m := Motor{}
//	m.Ride()
//
//	s := Sqlite{}
//	s.Create()
//
//}
//
//type Motor struct {
//
//}
//func (o *Motor) Ride() {}
//
//type Sqlite struct {
//
//}
//
//func (o *Sqlite) Create() {
//
//}
