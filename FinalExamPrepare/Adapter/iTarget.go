package main

type ITarget interface {
	start()
	stop()
	on()
	off()
}
