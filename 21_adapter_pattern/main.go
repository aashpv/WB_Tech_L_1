package main

import "fmt"

// OldDevice представляет собой старое устройство,
// которое имеет методы PlugIn и Unplug.
type OldDevice struct{}

// PlugIn подключает старое устройство.
func (o *OldDevice) PlugIn() {
	fmt.Println("Старое устройство подключено.")
}

// Unplug отключает старое устройство.
func (o *OldDevice) Unplug() {
	fmt.Println("Старое устройство отключено.")
}

// NewDevice - интерфейс, который представляет новое устройство,
// которое нужно зарядить.
type NewDevice interface {
	Charge()
}

// Adapter адаптирует старое устройство под новый интерфейс NewDevice.
type Adapter struct {
	oldDevice *OldDevice
}

// Charge реализует метод интерфейса NewDevice,
// используя старое устройство для подключения и отключения.
func (a *Adapter) Charge() {
	a.oldDevice.PlugIn() // Подключаем старое устройство
	fmt.Println("Зарядка производится.")
	a.oldDevice.Unplug() // Отключаем старое устройство
}

func main() {
	oldDevice := &OldDevice{}                 // Создаём экземпляр старого устройства
	adapter := &Adapter{oldDevice: oldDevice} // Создаём адаптер, который использует старое устройство

	var newDevice NewDevice = adapter // Присваиваем адаптер переменной типа NewDevice
	newDevice.Charge()                // Вызываем метод Charge, который использует старое устройство
}
