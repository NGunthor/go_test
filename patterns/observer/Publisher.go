package observer

//Publisher interface
type Publisher interface {
	Attach(observer Observer)
	Unpin(observer Observer)
	Notify()
	Show()
}
