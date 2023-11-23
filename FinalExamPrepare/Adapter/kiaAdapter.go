package main

type KiaAdapter struct {
	kia *Kia
}

func (k *KiaAdapter) start() {
	k.kia.turnOnEngine()
}

func (k *KiaAdapter) stop() {
	k.kia.turnOffEngine()
}

func (k *KiaAdapter) on() {
	k.kia.turnOnLights()
}

func (k *KiaAdapter) off() {
	k.kia.turnOffLights()
}
