package state

type State interface {
	addItem(int) error
	requestItem() error
	insertMonet(money int) error
	dispenseItem() error
}

//func addItem(count int) error {
//
//}
//
//func requestItem() error {
//
//}
//
//func insertMonet(money int) error {
//
//}
//
//func dispenseItem() error {
//
//}
