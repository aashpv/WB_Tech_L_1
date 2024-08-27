package main

import "fmt"

type OldDevice struct{}

func (o *OldDevice) PlugIn() {
	fmt.Println("Старое устройство подключено.")
}

func (o *OldDevice) Unplug() {
	fmt.Println("Старое устройство отключено.")
}
