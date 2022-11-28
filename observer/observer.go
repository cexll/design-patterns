package observer

type Observer interface {
	update(string)
	getId() string
}
