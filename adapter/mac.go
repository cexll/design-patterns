package adapter

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoLightingPort() {
	fmt.Println("lightning connector is plugged into mac machine.")
}
