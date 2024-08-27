package main

func main() {
	oldDevice := &OldDevice{}
	adapter := &Adapter{oldDevice: oldDevice}

	var newDevice NewDevice = adapter
	newDevice.Charge()
}
