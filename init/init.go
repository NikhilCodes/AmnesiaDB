package init

import (
	"amnesia-db/config"
	"fmt"
	"sync"
)

func All() {
	config.InitConfig()

	KeyStore()
	RandomSeed()
	go Socket()
	fmt.Println("LISTENING...")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
