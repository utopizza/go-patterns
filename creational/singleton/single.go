package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

var singleInstance *Single

type Single struct{}

func GetSingleInstance() *Single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now.")
				singleInstance = &Single{}
			})
	} else {
		fmt.Println("single instance already created.")
	}

	return singleInstance
}
