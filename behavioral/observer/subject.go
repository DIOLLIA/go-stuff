package main

type Subject interface {
	register(observer Observer)
	unregister(observer Observer)
	notifyAll()
}
