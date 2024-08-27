package main

import "fmt"

type Adapter struct {
	oldDevice *OldDevice
}

func (a *Adapter) Charge() {
	a.oldDevice.PlugIn() // Подключаем старое устройство
	fmt.Println("Зарядка производится.")
	a.oldDevice.Unplug() // Отключаем старое устройство
}
