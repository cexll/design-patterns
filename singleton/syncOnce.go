package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleOnce struct{}

var singleInstanceOnce *singleOnce

func getSingleInstance() *singleOnce {
	if singleInstanceOnce == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstanceOnce = &singleOnce{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstanceOnce
}
